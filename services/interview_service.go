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

	// Generate questions (this would be your AI logic)
	questions := s.generateQuestions(createdSession)

	// Return response
	response := &responses.InterviewSessionResponse{
		SessionID: createdSession.SessionID,
		Questions: questions,
	}

	return response, nil
}

// generateQuestions generates interview questions based on the session
func (s *InterviewService) generateQuestions(session *models.InterviewSession) []responses.Question {
	// This is where you'd implement your AI logic to generate questions
	// For now, returning mock questions
	questions := []responses.Question{
		{
			Question: "Tell me about a time when you had to work under pressure.",
			Hints:    []string{"Use the STAR method", "Be specific about the situation", "Highlight your actions and results"},
		},
		{
			Question: "Describe a situation where you had to resolve a conflict with a team member.",
			Hints:    []string{"Focus on communication", "Show how you found a solution", "Demonstrate leadership skills"},
		},
	}

	return questions
}
