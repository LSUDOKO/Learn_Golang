package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	mongoURI       = getEnv("MONGO_URI", "mongodb://localhost:27017")
	jwtSecret      = []byte(getEnv("JWT_SECRET", "change_this_secret"))
	googleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	googleSecret   = os.Getenv("GOOGLE_CLIENT_SECRET")
	googleRedirect = getEnv("GOOGLE_REDIRECT_URL", "http://localhost:9090/auth/google/callback")
	oauthConf      *oauth2.Config
	userCollection *mongo.Collection
	jwtExpiryHour  = 48
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash string             `bson:"password_hash,omitempty" json:"-"`
	GoogleID     string             `bson:"google_id,omitempty" json:"-"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to system env")
	}

	// Now reload env vars
	googleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	googleSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	googleRedirect = getEnv("GOOGLE_REDIRECT_URL", "http://localhost:9090/auth/google/callback")

	// init oauth config
	oauthConf = &oauth2.Config{
		ClientID:     googleClientID,
		ClientSecret: googleSecret,
		RedirectURL:  googleRedirect,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     google.Endpoint,
	}

	// connect to mongo
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("mongo new client:", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		log.Fatal("mongo connect:", err)
	}
	db := client.Database("authdb")
	userCollection = db.Collection("users")
	log.Println("Connected to MongoDB")

	// Gin router
	r := gin.Default()

	// Public routes
	r.POST("/signup", signupHandler)
	r.POST("/login", loginHandler)
	r.GET("/auth/google", googleLoginHandler)             // redirect to Google
	r.GET("/auth/google/callback", googleCallbackHandler) // callback

	// Protected example
	protected := r.Group("/api")
	protected.Use(jwtMiddleware())
	protected.GET("/me", meHandler)

	log.Println("Server running on :9090")
	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}

// ---------- Handlers ----------

// Signup: email + password
func signupHandler(c *gin.Context) {
	var body struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check existing user
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cnt, err := userCollection.CountDocuments(ctx, bson.M{"email": body.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to hash password"})
		return
	}

	user := User{
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
	}

	res, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	// do not return password hash
	user.PasswordHash = ""
	c.JSON(http.StatusCreated, gin.H{"user": user})
}

// Login: email + password -> returns JWT
func loginHandler(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	if err := userCollection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&user); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if user.PasswordHash == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "use Google login for this account"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	tokenString, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Google OAuth: redirect to Google's consent screen
func googleLoginHandler(c *gin.Context) {
	state := generateState() // optionally store server-side if you want CSRF protection
	url := oauthConf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// Google callback: exchange code -> get userinfo -> create/find user -> return JWT
func googleCallbackHandler(c *gin.Context) {
	ctx := context.Background()
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing code"})
		return
	}

	token, err := oauthConf.Exchange(ctx, code)
	if err != nil {
		log.Println("Token exchange error:", err) // log the actual error from Google
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Use token to call Google userinfo endpoint
	client := oauthConf.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil || resp.StatusCode != 200 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch userinfo"})
		return
	}
	defer resp.Body.Close()

	var info struct {
		ID            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locale"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse userinfo"})
		return
	}

	// Find or create user by Google ID or email
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err = userCollection.FindOne(ctx2, bson.M{"$or": []bson.M{{"google_id": info.ID}, {"email": info.Email}}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		// create
		newUser := User{
			Name:      info.Name,
			Email:     info.Email,
			GoogleID:  info.ID,
			CreatedAt: time.Now(),
		}
		res, err := userCollection.InsertOne(ctx2, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
			return
		}
		newUser.ID = res.InsertedID.(primitive.ObjectID)
		user = newUser
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	} else {
		// existing user found: ensure google_id is set
		if user.GoogleID == "" {
			_, _ = userCollection.UpdateOne(ctx2, bson.M{"_id": user.ID}, bson.M{"$set": bson.M{"google_id": info.ID}})
		}
	}

	// Generate JWT & return (or redirect)
	tokenString, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	// You can redirect with token in query, or return JSON. We'll return JSON.
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Protected route example
func meHandler(c *gin.Context) {
	claimUser, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no user in context"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"me": claimUser})
}

// ---------- Helpers ----------

func generateJWT(user User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"email":   user.Email,
		"name":    user.Name,
		"exp":     time.Now().Add(time.Hour * time.Duration(jwtExpiryHour)).Unix(),
		"iat":     time.Now().Unix(),
		"iss":     "myapp",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header missing"})
			return
		}
		// Expect "Bearer <token>"
		var tokenString string
		fmt.Sscanf(auth, "Bearer %s", &tokenString)
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// verify alg
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}

		// put user info in context
		c.Set("user", claims)
		c.Next()
	}
}

func generateState() string {
	// For demonstration return static; in production create random and store to validate
	return "state-token"
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
