package repositories

import (
	"context"
	"errors"
	"math/rand"
	"stormhacks-be/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InterviewRepository handles MongoDB operations for interview sessions and related data
type InterviewRepository struct {
	sessionsCollection     *mongo.Collection
	questionsCollection    *mongo.Collection
}

// NewInterviewRepository creates a new interview repository with all collections
func NewInterviewRepository(db *mongo.Database) *InterviewRepository {
	return &InterviewRepository{
		sessionsCollection:  db.Collection("interview_sessions"),
		questionsCollection: db.Collection("question_bank"),
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
	result, err := r.sessionsCollection.InsertOne(ctx, session)
	if err != nil {
		return nil, err
	}

	// Set the generated ID
	session.ID = result.InsertedID.(primitive.ObjectID)
	return session, nil
}

// GetBySessionID retrieves an interview session by session ID
func (r *InterviewRepository) GetBySessionID(sessionID string) (*models.InterviewSession, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var session models.InterviewSession
	err := r.sessionsCollection.FindOne(ctx, bson.M{"session_id": sessionID}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("not found")
		}
		return nil, err
	}

	return &session, nil
}



// GetQuestionsByBehavioralTopic retrieves questions by behavioral topic
func (r *InterviewRepository) GetQuestionsByBehavioralTopic(topic string) ([]models.QuestionBank, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.questionsCollection.Find(ctx, bson.M{"behavioralTopic": topic})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var questions []models.QuestionBank
	if err = cursor.All(ctx, &questions); err != nil {
		return nil, err
	}

	return questions, nil
}

// GetRandomQuestionsByTopics retrieves random questions for given topics
func (r *InterviewRepository) GetRandomQuestionsByTopics(topics []string) ([]models.QuestionBank, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pad or trim topics to 3
	for len(topics) < 3 {
		topics = append(topics, "General")
	}
	if len(topics) > 3 {
		topics = topics[:3]
	}

	rand.Seed(time.Now().UnixNano())
	var selected []models.QuestionBank

	for _, topic := range topics {
		filter := bson.M{"behavioralTopic": topic}
		cursor, err := r.questionsCollection.Find(ctx, filter)
		if err != nil {
			return nil, err
		}
		defer cursor.Close(ctx)

		var questions []models.QuestionBank
		if err := cursor.All(ctx, &questions); err != nil {
			return nil, err
		}

		if len(questions) == 0 {
			continue
		}

		// Pick 1 random question from this topic
		q := questions[rand.Intn(len(questions))]
		selected = append(selected, q)
	}

	return selected, nil
}


