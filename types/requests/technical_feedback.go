package requests

type TechnicalFeedbackInput struct {
	SessionID    string `json:"sessionId" validate:"required"`
	QuestionID   string `json:"questionId" validate:"required"`
	UserCode     string `json:"userCode" validate:"required"`
	HintsUsed    int    `json:"hintsUsed" validate:"required"`
	IsCompleted  bool   `json:"isCompleted" validate:"required"`
	TimeTaken    int    `json:"timeTaken" validate:"required"` // in seconds
}