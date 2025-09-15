package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// manager struct to hold connection details
type manager struct {
	Connection *mongo.Client
	Ctx        context.Context
	Cancel     context.CancelFunc
}

var mgr manager

// Connect to MongoDB
func connectDb() {
	uri := "mongodb://localhost:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("Error creating mongo client:", err)
		cancel()
		return
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println("Could not ping MongoDB:", err)
		cancel()
		return
	}

	mgr = manager{
		Connection: client,
		Ctx:        ctx,
		Cancel:     cancel,
	}

	fmt.Println("Connected to MongoDB successfully")
}

// Close the connection
func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func main() {
	connectDb()
	defer Close(mgr.Connection, mgr.Ctx, mgr.Cancel)
}
