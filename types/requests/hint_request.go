package requests

// HintRequest represents the input for generating hints
type HintRequest struct {
	SessionID     string   `json:"sessionId" validate:"required"`
	Question      string   `json:"question" validate:"required"`
	PreviousHints []string `json:"previousHints,omitempty"`
	UserCode      string   `json:"userCode" validate:"required"`
	UserSpeech    string   `json:"userSpeech" validate:"required"` // What the user actually said (speech-to-text)
}
