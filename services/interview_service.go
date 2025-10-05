package services

import (
	"errors"
	"log"
	"stormhacks-be/models"
	"stormhacks-be/repositories"
	"stormhacks-be/types/enums"
	"stormhacks-be/types/requests"
	"stormhacks-be/types/responses"
	"strconv"
)

// InterviewService handles interview business logic
type InterviewService struct {
	interviewRepo *repositories.InterviewRepository
}

// NewInterviewService creates a new interview service
func NewInterviewService(interviewRepo *repositories.InterviewRepository) *InterviewService {
	return &InterviewService{
		interviewRepo: interviewRepo,
	}
}

func (s *InterviewService) GetInterviewSession(sessionID int) (*models.InterviewSession, error) {
	return s.interviewRepo.GetBySessionID(sessionID)
}

// CreateInterviewSession creates a new interview session
func (s *InterviewService) CreateInterviewSession(input requests.InterviewSessionInput) (*responses.InterviewSessionResponse, error) {
	// Check if session already exists
	existingSession, err := s.interviewRepo.GetBySessionID(input.SessionID)
	if err != nil && err.Error() != "not found" {
		return nil, err
	}
	if existingSession != nil {
		return nil, errors.New("session already exists")
	}
	// Create interview session model
	session := &models.InterviewSession{
		SessionID:        input.SessionID,
		ParsedResumeText: input.ParsedResumeText,
		JobTitle:         input.JobTitle,
		JobInfo:          input.JobInfo,
		CompanyName:        input.CompanyName,
		AdditionalInfo:     input.AdditionalInfo,
		InterviewType:      input.TypeOfInterview,
		BehaviouralTopics:  input.BehaviouralTopics,
		TechnicalDifficulty: input.TechnicalDifficulty,
	}

	// Save to database
	createdSession, err := s.interviewRepo.Create(session)
	if err != nil {
		return nil, err
	}

	// Return response
	response := &responses.InterviewSessionResponse{
		SessionID: createdSession.SessionID,
	}

	return response, nil
}

// GenerateInterviewQuestions generates interview questions based on the session
func (s *InterviewService) GenerateInterviewQuestions(sessionID int) (*responses.InterviewSessionQuestionsResponse, error) {
	// Get the session first
	session, err := s.interviewRepo.GetBySessionID(sessionID)
	if err != nil {
		return nil, err
	}

	// Convert enum topics to strings for database query
	topicStrings := make([]string, len(session.BehaviouralTopics))
	for i, topic := range session.BehaviouralTopics {
		topicStrings[i] = string(topic)
	}

	// Get random questions based on behavioral topics
	questions, err := s.interviewRepo.GetRandomQuestionsByTopics(topicStrings)
	if err != nil {
		return nil, err
	}

	// Use Gemini to customize questions based on job description and resume
	googleGeminiService := NewGoogleGeminiService()
	customizedQuestions, err := googleGeminiService.CustomizeInterviewQuestions(session, questions)
	if err != nil {
		// If Gemini fails, fall back to original questions
		log.Printf("Warning: Failed to customize questions with Gemini: %v. Using original questions.", err)
	}

	// Convert to response format
	var responseQuestions []responses.InterviewQuestion
	for i, q := range customizedQuestions {
		responseQuestions = append(responseQuestions, responses.InterviewQuestion{
			ID:       "q" + strconv.Itoa(i+1),
			Topic:    string(q.BehavioralTopic),
			Question: q.Question,
			Hints:    q.Hints, // Use AI-generated hints
		})
	}

	return &responses.InterviewSessionQuestionsResponse{
		SessionID: sessionID,
		Questions: responseQuestions,
	}, nil
}

// GetAllBehavioralTopics returns all available behavioral topics
func (s *InterviewService) GetAllBehavioralTopics() []string {
	topics := enums.GetAllBehaviouralTopics()
	topicStrings := make([]string, len(topics))
	for i, topic := range topics {
		topicStrings[i] = string(topic)
	}
	return topicStrings
}

// GetAvailableInterviewTypes returns all available interview types as strings
func (s *InterviewService) GetAvailableInterviewTypes() []string {
	return []string{"technical", "behavioral", "both"}
}

// GetQuestionsByTopic returns questions for a specific behavioral topic
func (s *InterviewService) GetQuestionsByTopic(topic string) ([]models.QuestionBank, error) {
	return s.interviewRepo.GetQuestionsByBehavioralTopic(topic)
}

// generateHintsForTopic generates appropriate hints based on the behavioral topic
func generateHintsForTopic(topic enums.BehaviouralTopic) []string {
	switch topic {
	case enums.BehaviouralTopicGeneral:
		return []string{"Be authentic", "Highlight relevant experience", "Show enthusiasm for the role"}
	case enums.BehaviouralTopicWorkplaceBehavior:
		return []string{"Focus on professionalism", "Show cultural awareness", "Demonstrate teamwork"}
	case enums.BehaviouralTopicLeadership:
		return []string{"Use STAR method", "Show leadership qualities", "Highlight team impact"}
	case enums.BehaviouralTopicProblemSolving:
		return []string{"Show analytical thinking", "Explain your process", "Highlight the outcome"}
	case enums.BehaviouralTopicConflictResolution:
		return []string{"Focus on communication", "Show empathy", "Highlight resolution skills"}
	case enums.BehaviouralTopicAdaptability:
		return []string{"Highlight flexibility", "Focus on results", "Show learning ability"}
	case enums.BehaviouralTopicTimeManagement:
		return []string{"Show organization skills", "Highlight prioritization", "Demonstrate efficiency"}
	case enums.BehaviouralTopicCustomerFocus:
		return []string{"Focus on customer needs", "Show problem-solving", "Highlight satisfaction"}
	case enums.BehaviouralTopicInnovationCreativity:
		return []string{"Show creativity", "Highlight innovation", "Focus on results"}
	default:
		return []string{"Use STAR method", "Be specific", "Show your impact"}
	}
}


func (s *InterviewService) GenerateInterviewFeedback(input requests.InterviewFeedbackInput) (*responses.InterviewFeedbackResponse, error) {
	existingSession, err := s.interviewRepo.GetBySessionID(input.SessionID)
	if err != nil && err.Error() != "not found" {
		return nil, err
	}
	if existingSession == nil {
		return nil, errors.New("session not found")
	}
	
	// Create Gemini service and generate feedback
	googleGeminiService := NewGoogleGeminiService()
	feedbackResponse, err := googleGeminiService.GenerateInterviewFeedback(existingSession, input.InterviewQuestionsWithAnswers)
	if err != nil {
		return nil, err
	}
	
	return feedbackResponse, nil
}