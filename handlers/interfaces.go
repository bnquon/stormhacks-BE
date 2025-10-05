package handlers

import (
	"stormhacks-be/models"
	"stormhacks-be/types/requests"
	"stormhacks-be/types/responses"
)

// InterviewServiceInterface defines the interface for interview service
type InterviewServiceInterface interface {
	CreateInterviewSession(input requests.InterviewSessionInput) (*responses.InterviewSessionResponse, error)
	GenerateInterviewQuestions(sessionID string) (*responses.InterviewSessionQuestionsResponse, error)
	GenerateInterviewFeedback(input requests.InterviewFeedbackInput) (*responses.InterviewFeedbackResponse, error)
	GetTechnicalQuestion(difficulty string) (*models.TechnicalBank, error)
}
