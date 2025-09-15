package main

//this context contain all informaton about the request that handle might be needed
//context contain cokkies handlers
//json {} [{},{}]
//XML <tag></tag>
/*func test(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "pong",
	})
}
Here, c *gin.Context is passed to every handler.
It’s like a toolbox for handling the request and response.

What gin.Context does
It provides methods to:

Access Request Data

Get query parameters:

go
Copy code
name := c.Query("name") // /ping?name=Arpit
Get path parameters:

go
Copy code
r.GET("/user/:id", func(c *gin.Context) {
    id := c.Param("id") // /user/10 → id = "10"
})
Get POST form data:

go
Copy code
email := c.PostForm("email")
Get JSON body:

go
Copy code
var user User
if err := c.BindJSON(&user); err == nil {
    // use user struct
}
Send Responses

Send JSON:

go
Copy code
c.JSON(200, gin.H{"status": "ok"})
Send string:

go
Copy code
c.String(200, "Hello World")
Send HTML:

go
Copy code
c.HTML(200, "index.tmpl", gin.H{"title": "Home"})
Control Request Flow

You can abort a request if validation fails:

go
Copy code
if !authenticated {
    c.AbortWithStatusJSON(401, gin.H{"error": "unauthorized"})
    return
}
Share Data Across Middleware

Store values during request handling:

go
Copy code
c.Set("userID", 123)
Retrieve it later in another middleware or handler:

go
Copy code
id, exists := c.Get("userID")*/
// import {
// 	"github.com/gin-gonic/gin"
// }
