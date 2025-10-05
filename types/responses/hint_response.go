package responses

// HintResponse represents the response for hint generation
type HintResponse struct {
	SessionID         string `json:"sessionId"`
	ConversationalHint string `json:"conversationalHint"` // For text-to-speech
	HintSummary       string `json:"hintSummary"`         // For display
}
