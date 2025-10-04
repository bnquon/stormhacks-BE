package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	globalClient   *Client
	globalDatabase *mongo.Database
)

// Client wraps the MongoDB client
type Client struct {
	*mongo.Client
	Database *mongo.Database
}

// Config holds MongoDB configuration
type Config struct {
	URI      string
	Database string
	Username string
	Password string
}

// NewMongoClient creates a new MongoDB client
func NewMongoClient(config Config) (*Client, error) {
	// Set client options with Server API
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(config.URI).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Test the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	// Get database
	db := client.Database(config.Database)

	return &Client{
		Client:   client,
		Database: db,
	}, nil
}

// DefaultConfig returns a default MongoDB configuration from environment variables
func DefaultConfig() Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		uri = "mongodb://localhost:27017"
	}

	database := os.Getenv("MONGODB_DATABASE")
	if database == "" {
		database = "stormhacks"
	}

	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")

	return Config{
		URI:      uri,
		Database: database,
		Username: username,
		Password: password,
	}
}

// GetCollection returns a MongoDB collection
func (c *Client) GetCollection(name string) *mongo.Collection {
	return c.Database.Collection(name)
}

// InitDatabase initializes the global database connection
func InitDatabase() error {
	config := DefaultConfig()
	client, err := NewMongoClient(config)
	if err != nil {
		return err
	}
	
	globalClient = client
	globalDatabase = client.Database
	return nil
}

// GetDatabase returns the global database instance
func GetDatabase() *mongo.Database {
	return globalDatabase
}

// CloseDatabase closes the global database connection
func CloseDatabase() error {
	if globalClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return globalClient.Disconnect(ctx)
	}
	return nil
}
