package repositories

import (
	"context"
	"errors"
	"time"
	"stormhacks-be/types/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InterviewRepository handles MongoDB operations for interview sessions
type InterviewRepository struct {
	collection *mongo.Collection
}

// NewInterviewRepository creates a new interview repository
func NewInterviewRepository(collection *mongo.Collection) *InterviewRepository {
	return &InterviewRepository{
		collection: collection,
	}
}

// Create creates a new interview session in MongoDB
func (r *InterviewRepository) Create(session *models.InterviewSession) (*models.InterviewSession, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set timestamps
	now := time.Now()
	session.CreatedAt = now

	// Insert document
	result, err := r.collection.InsertOne(ctx, session)
	if err != nil {
		return nil, err
	}

	// Set the generated ID
	session.ID = result.InsertedID.(primitive.ObjectID)
	return session, nil
}

// GetBySessionID retrieves an interview session by session ID
func (r *InterviewRepository) GetBySessionID(sessionID int) (*models.InterviewSession, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var session models.InterviewSession
	err := r.collection.FindOne(ctx, bson.M{"session_id": sessionID}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return &session, nil
}
