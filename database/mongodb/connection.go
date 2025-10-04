package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// Set client options
	clientOptions := options.Client().ApplyURI(config.URI)

	if config.Username != "" && config.Password != "" {
		clientOptions.SetAuth(options.Credential{
			Username: config.Username,
			Password: config.Password,
		})
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Test the connection
	err = client.Ping(ctx, nil)
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

// DefaultConfig returns a default MongoDB configuration
func DefaultConfig() Config {
	return Config{
		URI:      "mongodb://localhost:27017",
		Database: "stormhacks",
		Username: "",
		Password: "",
	}
}

// GetCollection returns a MongoDB collection
func (c *Client) GetCollection(name string) *mongo.Collection {
	return c.Database.Collection(name)
}
