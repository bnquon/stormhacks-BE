package migrations

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RunMigrations runs all database migrations
func RunMigrations(db *mongo.Database) error {
	log.Println("Starting database migrations...")

	// Create indexes for all collections
	if err := createIndexes(db); err != nil {
		return fmt.Errorf("failed to create indexes: %v", err)
	}

	// Seed question bank
	if err := seedQuestionBank(db); err != nil {
		log.Printf("Warning: Failed to seed question bank: %v", err)
		log.Println("Continuing without seeding...")
	}

	log.Println("Database migrations completed successfully!")
	return nil
}

// createIndexes creates necessary indexes for all collections
func createIndexes(db *mongo.Database) error {
	ctx := context.Background()

	// Indexes for interview_sessions
	interviewSessionsCollection := db.Collection("interview_sessions")
	interviewSessionIndexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "session_id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "created_at", Value: -1}},
		},
		{
			Keys: bson.D{{Key: "job_title", Value: 1}},
		},
	}
	_, err := interviewSessionsCollection.Indexes().CreateMany(ctx, interviewSessionIndexes)
	if err != nil {
		return fmt.Errorf("failed to create interview_sessions indexes: %v", err)
	}

	// Indexes for question_bank
	questionsCollection := db.Collection("question_bank")
	questionIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "behavioralTopic", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "difficulty", Value: 1}},
		},
	}
	_, err = questionsCollection.Indexes().CreateMany(ctx, questionIndexes)
	if err != nil {
		return fmt.Errorf("failed to create question_bank indexes: %v", err)
	}



	log.Println("All indexes created successfully!")
	return nil
}

// seedQuestionBank seeds the question bank with sample questions
func seedQuestionBank(db *mongo.Database) error {
	ctx := context.Background()
	collection := db.Collection("question_bank")

	// Check if questions already exist
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("failed to count questions: %v", err)
	}

	if count > 0 {
		log.Println("Questions already exist, skipping seed...")
		return nil
	}

	// Sample questions for seeding
	var documents []interface{}
	questions := []bson.M{
		// General
		{"behavioralTopic": "General", "question": "Tell me about yourself and how your background prepared you for this role."},
		{"behavioralTopic": "General", "question": "What motivates you the most at work?"},
		{"behavioralTopic": "General", "question": "Describe your proudest professional achievement."},
		{"behavioralTopic": "General", "question": "Tell me about a time you made a mistake and what you learned."},
		{"behavioralTopic": "General", "question": "What do you consider your biggest strength and how has it helped you succeed?"},

		// Workplace Behavior
		{"behavioralTopic": "Workplace Behavior", "question": "Tell me about a time you had to adapt to a new workplace culture."},
		{"behavioralTopic": "Workplace Behavior", "question": "Describe how you handle differences in work styles with coworkers."},
		{"behavioralTopic": "Workplace Behavior", "question": "Give an example of how you helped create a positive team environment."},
		{"behavioralTopic": "Workplace Behavior", "question": "Tell me about a time you went above and beyond at work."},
		{"behavioralTopic": "Workplace Behavior", "question": "Describe how you maintain professionalism under pressure."},

		// Leadership
		{"behavioralTopic": "Leadership", "question": "Tell me about a time you led a team through a difficult challenge."},
		{"behavioralTopic": "Leadership", "question": "Describe a situation where you inspired others to improve their performance."},
		{"behavioralTopic": "Leadership", "question": "Tell me about a time you had to make a tough decision as a leader."},
		{"behavioralTopic": "Leadership", "question": "Give an example of how you delegated tasks effectively."},
		{"behavioralTopic": "Leadership", "question": "Describe a time you developed or mentored a teammate."},

		// Problem Solving
		{"behavioralTopic": "Problem Solving", "question": "Tell me about a complex problem you solved at work."},
		{"behavioralTopic": "Problem Solving", "question": "Describe a time you found a creative solution to a difficult issue."},
		{"behavioralTopic": "Problem Solving", "question": "Tell me about a situation where you used data to make a decision."},
		{"behavioralTopic": "Problem Solving", "question": "Describe a time you solved a problem under tight deadlines."},
		{"behavioralTopic": "Problem Solving", "question": "Give an example of a process you improved through analysis or innovation."},

		// Conflict Resolution
		{"behavioralTopic": "Conflict Resolution", "question": "Tell me about a time you resolved a disagreement with a coworker."},
		{"behavioralTopic": "Conflict Resolution", "question": "Describe how you handled a situation with a difficult team member."},
		{"behavioralTopic": "Conflict Resolution", "question": "Tell me about a time you mediated between two parties in conflict."},
		{"behavioralTopic": "Conflict Resolution", "question": "Give an example of when you gave constructive feedback that led to improvement."},
		{"behavioralTopic": "Conflict Resolution", "question": "Describe a time you disagreed with your manager and how you handled it."},

		// Adaptability
		{"behavioralTopic": "Adaptability", "question": "Tell me about a time you had to adapt quickly to a major change."},
		{"behavioralTopic": "Adaptability", "question": "Describe a situation where you had to learn something new to complete a project."},
		{"behavioralTopic": "Adaptability", "question": "Give an example of how you handled multiple priorities at once."},
		{"behavioralTopic": "Adaptability", "question": "Tell me about a time you succeeded despite unexpected obstacles."},
		{"behavioralTopic": "Adaptability", "question": "Describe how you adapt your communication style to different audiences."},

		// Time Management
		{"behavioralTopic": "Time Management", "question": "Tell me about a time you had to meet a very tight deadline."},
		{"behavioralTopic": "Time Management", "question": "Describe how you plan and organize your daily tasks."},
		{"behavioralTopic": "Time Management", "question": "Tell me about a time you missed a deadline and what you learned."},
		{"behavioralTopic": "Time Management", "question": "Give an example of how you manage conflicting priorities."},
		{"behavioralTopic": "Time Management", "question": "Describe a tool or technique you use to stay on schedule."},

		// Customer Focus
		{"behavioralTopic": "Customer Focus", "question": "Tell me about a time you went above and beyond for a customer."},
		{"behavioralTopic": "Customer Focus", "question": "Describe a situation where you handled a difficult client."},
		{"behavioralTopic": "Customer Focus", "question": "Tell me about how you gathered and used customer feedback."},
		{"behavioralTopic": "Customer Focus", "question": "Give an example of when you improved customer satisfaction."},
		{"behavioralTopic": "Customer Focus", "question": "Describe a time you balanced customer needs with company policy."},

		// Innovation & Creativity
		{"behavioralTopic": "Innovation & Creativity", "question": "Tell me about a time you introduced a new idea that improved a process."},
		{"behavioralTopic": "Innovation & Creativity", "question": "Describe a time you thought outside the box to solve a problem."},
		{"behavioralTopic": "Innovation & Creativity", "question": "Give an example of how you encouraged creativity in your team."},
		{"behavioralTopic": "Innovation & Creativity", "question": "Tell me about a time you challenged an existing approach to find a better solution."},
		{"behavioralTopic": "Innovation & Creativity", "question": "Describe a creative solution you implemented successfully."},
	}

	// Convert to interface{} slice
	for _, question := range questions {
		documents = append(documents, question)
	}

	// Insert questions
	_, err = collection.InsertMany(ctx, documents)
	if err != nil {
		return fmt.Errorf("failed to insert questions: %v", err)
	}

	log.Printf("Seeded %d questions", len(documents))
	return nil
}