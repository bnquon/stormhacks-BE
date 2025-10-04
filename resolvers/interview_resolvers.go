package resolvers

import (
	"context"
	"errors"
	"fmt"
	"stormhacks-be/database/mongodb"
	"stormhacks-be/models"
	enums "stormhacks-be/utils/enums"
	"time"

	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InterviewResolvers contains all interview-related resolvers
type InterviewResolvers struct {
	db *mongo.Database
}

// NewInterviewResolvers creates a new instance of InterviewResolvers
func NewInterviewResolvers() *InterviewResolvers {
	return &InterviewResolvers{
		db: mongodb.GetDatabase(),
	}
}

// CreateInterviewSession resolves the createInterviewSession mutation
func (ir *InterviewResolvers) CreateInterviewSession(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid input provided")
	}

	// Convert input to InterviewSession model
	session := models.InterviewSession{
		SessionID:         int(input["sessionId"].(int)),
		ParsedResumeText:  input["parsedResumeText"].(string),
		JobTitle:          input["jobTitle"].(string),
		JobInfo:           input["jobInfo"].(string),
		BehaviouralTopics: []string{string(enums.General)}, // Default value
		CreatedAt:         time.Now(),
	}
	// Handle optional fields
	if companyName, ok := input["companyName"].(string); ok && companyName != "" {
		session.CompanyName = &companyName
	}
	if additionalInfo, ok := input["additionalInfo"].(string); ok && additionalInfo != "" {
		session.AdditionalInfo = &additionalInfo
	}
	if typeOfInterview, ok := input["typeOfInterview"].(string); ok && typeOfInterview != "" {
		session.TypeOfInterview = &typeOfInterview
	}
	if technicalDifficulty, ok := input["technicalDifficulty"].(string); ok && technicalDifficulty != "" {
		session.TechnicalDifficulty = &technicalDifficulty
	}
	if topics, ok := input["behaviouralTopics"].([]interface{}); ok && len(topics) > 0 {
		behaviouralTopics := make([]string, len(topics))
		for i, topic := range topics {
			behaviouralTopics[i] = topic.(string)
		}
		session.BehaviouralTopics = behaviouralTopics
	}

	// Insert into database
	collection := ir.db.Collection("interview_sessions")
	result, err := collection.InsertOne(context.Background(), session)
	if err != nil {
		return nil, fmt.Errorf("failed to create interview session: %v", err)
	}

	// Get the created session
	session.ID = result.InsertedID.(primitive.ObjectID)
	return session, nil
}

// GetInterviewSession resolves the getInterviewSession query
func (ir *InterviewResolvers) GetInterviewSession(params graphql.ResolveParams) (interface{}, error) {
	sessionID, ok := params.Args["sessionId"].(int)
	if !ok {
		return nil, errors.New("sessionId is required")
	}

	collection := ir.db.Collection("interview_sessions")
	var session models.InterviewSession
	
	err := collection.FindOne(context.Background(), bson.M{"session_id": sessionID}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("interview session not found")
		}
		return nil, fmt.Errorf("failed to get interview session: %v", err)
	}

	return session, nil
}

// GenerateInterviewQuestions resolves the generateInterviewQuestions query
func (ir *InterviewResolvers) GenerateInterviewQuestions(params graphql.ResolveParams) (interface{}, error) {
	sessionID, ok := params.Args["sessionId"].(int)
	if !ok {
		return nil, errors.New("sessionId is required")
	}

	// Get the session first
	sessionInterface, err := ir.GetInterviewSession(params)
	if err != nil {
		return nil, err
	}

	session, ok := sessionInterface.(models.InterviewSession)
	if !ok {
		return nil, errors.New("failed to parse session data")
	}

	// Mock implementation - in real app, this would use AI to generate questions
	questions := []map[string]interface{}{
		{
			"question": fmt.Sprintf("Tell me about a time when you had to handle a difficult situation related to %s.", session.JobTitle),
			"hints":    []string{"Use the STAR method", "Be specific about your actions", "Focus on the outcome"},
		},
		{
			"question": "Describe a situation where you had to work with a difficult team member.",
			"hints":    []string{"Stay professional", "Focus on resolution", "Show emotional intelligence"},
		},
		{
			"question": "Give me an example of a time you failed and what you learned from it.",
			"hints":    []string{"Be honest about the failure", "Emphasize the learning", "Show growth mindset"},
		},
	}

	return map[string]interface{}{
		"sessionId": sessionID,
		"questions": questions,
	}, nil
}

// SubmitBehaviouralFeedback resolves the submitBehaviouralFeedback mutation
func (ir *InterviewResolvers) SubmitBehaviouralFeedback(params graphql.ResolveParams) (interface{}, error) {
	input, ok := params.Args["input"].(map[string]interface{})
	if !ok {
		return nil, errors.New("invalid input provided")
	}

	sessionID := int(input["sessionId"].(int))
	responsesInterface, ok := input["responses"].([]interface{})
	if !ok {
		return nil, errors.New("responses are required")
	}

	// Convert responses to proper format
	responses := make([]map[string]interface{}, len(responsesInterface))
	for i, resp := range responsesInterface {
		responseMap, ok := resp.(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid response format")
		}
		responses[i] = responseMap
	}

	// Mock implementation - in real app, this would use AI to analyze responses
	questionFeedback := make([]map[string]interface{}, len(responses))
	for i, response := range responses {
		question := response["question"].(string)
		responseText := response["response"].(string)

		// Simple scoring based on response length (mock implementation)
		score := 5
		if len(responseText) > 100 {
			score = 8
		} else if len(responseText) > 50 {
			score = 6
		}

		feedback := "Good response. "
		if score < 6 {
			feedback += "Consider providing more specific examples and details."
		} else {
			feedback += "Well-structured answer with good examples."
		}

		suggestions := []string{
			"Use the STAR method (Situation, Task, Action, Result)",
			"Provide specific examples with measurable outcomes",
			"Show how your actions impacted the situation",
		}

		questionFeedback[i] = map[string]interface{}{
			"question":    question,
			"response":    responseText,
			"score":       score,
			"feedback":    feedback,
			"suggestions": suggestions,
		}
	}

	return map[string]interface{}{
		"sessionId":        sessionID,
		"questionFeedback": questionFeedback,
	}, nil
}

// GetBehaviouralFeedback resolves the getBehaviouralFeedback query
func (ir *InterviewResolvers) GetBehaviouralFeedback(params graphql.ResolveParams) (interface{}, error) {
	sessionID, ok := params.Args["sessionId"].(int)
	if !ok {
		return nil, errors.New("sessionId is required")
	}

	// In a real implementation, you would fetch this from the database
	// For now, return a mock response
	return map[string]interface{}{
		"sessionId": sessionID,
		"questionFeedback": []map[string]interface{}{
			{
				"question":    "Sample question",
				"response":    "Sample response",
				"score":       7,
				"feedback":    "Good response with room for improvement",
				"suggestions": []string{"Provide more specific examples", "Use the STAR method"},
			},
		},
	}, nil
}
