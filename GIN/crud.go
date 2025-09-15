package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// User struct for demo
type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"` //use random id if we use omitempty
	Name  string             `bson:"name"`
	Email string             `bson:"email"`
}

// manager struct to hold DB connection
type manager struct {
	Connection *mongo.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
	collection *mongo.Collection
}

// Manager interface with CRUD methods
type Manager interface {
	Insert(interface{}) error
	GetAll() ([]User, error)
	DeleteData(primitive.ObjectID) error
}

// Connect to MongoDB
func connectDb() (*manager, error) {
	uri := "mongodb://localhost:27017"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		cancel()
		return nil, err
	}

	m := &manager{
		Connection: client,
		Ctx:        ctx,
		Cancel:     cancel,
		collection: client.Database("testdb").Collection("users"),
	}
	fmt.Println("Connected Suceesfully")
	return m, nil
}

// Insert new user
func (m *manager) Insert(data interface{}) error {
	_, err := m.collection.InsertOne(m.Ctx, data)
	return err
}

// Get all users
func (m *manager) GetAll() ([]User, error) {
	var users []User
	cursor, err := m.collection.Find(m.Ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(m.Ctx)

	for cursor.Next(m.Ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Delete user by ID
func (m *manager) DeleteData(id primitive.ObjectID) error {
	res, err := m.collection.DeleteOne(m.Ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		fmt.Println("Warning: No document was deleted")
	}
	return nil
}

// main function for testing
func main() {
	mgr, err := connectDb()
	if err != nil {
		fmt.Println("Error connecting DB:", err)
		return
	}
	defer mgr.Cancel()

	// Insert
	newUser := User{Name: "Arpit", Email: "arpit@example.com"}
	if err := mgr.Insert(newUser); err != nil {
		fmt.Println("Insert error:", err)
	}

	// GetAll
	users, err := mgr.GetAll()
	if err != nil {
		fmt.Println("GetAll error:", err)
	}
	fmt.Println("Users:", users)

	fmt.Printf("Type of ID: %T\n", users[0].ID)

	// Delete first user if exists
	if len(users) > 0 {
		err = mgr.DeleteData(users[0].ID)
		if err != nil {
			fmt.Println("Delete error:", err)
		} else {
			fmt.Println("Deleted user:", users[0].ID)
		}

		count, _ := mgr.collection.CountDocuments(mgr.Ctx, bson.D{})
		fmt.Println("Documents remaining:", count)
	}

}
