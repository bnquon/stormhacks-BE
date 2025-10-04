package responses

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
