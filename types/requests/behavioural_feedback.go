package requests

// BehaviouralFeedbackInput represents the input for submitting behavioural feedback
type BehaviouralFeedbackInput struct {
	SessionID int                `json:"sessionId" validate:"required"`
	Responses []QuestionResponse `json:"responses" validate:"required,min=1"`
}

// QuestionResponse represents a single question and its response
type QuestionResponse struct {
	Question string `json:"question" validate:"required"`
	Response string `json:"response" validate:"required"` // Extracted transcript for each question from FE
}