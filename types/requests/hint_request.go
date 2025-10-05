package requests

// HintRequest represents the input for generating hints
type HintRequest struct {
	SessionID     string   `json:"sessionId" validate:"required"`
	QuestionID    string   `json:"questionId" validate:"required"`
	PreviousHints []string `json:"previousHints,omitempty"`
	UserCode      string   `json:"userCode" validate:"required"`
	UserSpeech    string   `json:"userSpeech" validate:"required"` // What the user actually said (speech-to-text)
}
