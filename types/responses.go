package types

// Question represents a single interview question
type Question struct {
	Question string   `json:"question"`
	Hints    []string `json:"hints"`
}

// InterviewSessionResponse represents the response for an interview session
type InterviewSessionResponse struct {
	SessionID int        `json:"sessionId"`
	Questions []Question `json:"questions"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Details string `json:"details,omitempty"`
}
