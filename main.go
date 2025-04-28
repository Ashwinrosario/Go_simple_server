package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"net/http"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"simpleServer/routes"
)

var client *mongo.Client
var collection *mongo.Collection

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read MongoDB URI and other parameters from the .env file
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	port := os.Getenv("SERVER_PORT")

	// Check if all necessary environment variables are set
	if mongoURI == "" || dbName == "" || collectionName == "" || port == "" {
		log.Fatalf("Missing required environment variables. Check your .env file.")
	}

	// Create MongoDB client with the URI
	client, err = mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	// Connect to MongoDB with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to connect to the MongoDB cluster
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Test the connection by pinging the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Successfully connected to MongoDB!")

	// Initialize the collection
	collection = client.Database(dbName).Collection(collectionName)

	// Initialize routes
	http.HandleFunc("/users", routes.UsersHandler(collection))
	http.HandleFunc("/users/", routes.SpecificUserHandler(collection))

	// Start the server
	fmt.Printf("Server is running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
