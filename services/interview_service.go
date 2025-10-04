package services

import (
	"errors"
	"stormhacks-be/types/models"
	"stormhacks-be/types/requests"
	"stormhacks-be/types/responses"
	"stormhacks-be/repositories"
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
		SessionID:          input.SessionID,
		ParsedResumeText:  input.ParsedResumeText,
		JobTitle:           input.JobTitle,
		JobInfo:            input.JobInfo,
		CompanyName:        input.CompanyName,
		AdditionalInfo:     input.AdditionalInfo,
		TypeOfInterview:    input.TypeOfInterview,
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