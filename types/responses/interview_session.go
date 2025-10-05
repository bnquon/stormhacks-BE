package responses

import "stormhacks-be/models"

// InterviewSessionResponse represents the response for an interview session
type InterviewSessionResponse struct {
	SessionID int `json:"sessionId"`
}

// InterviewQuestion represents a single interview question
type InterviewQuestion struct {
	ID       string   `json:"id"`
	Topic    string   `json:"topic"`
	Question string   `json:"question"`
	Hints    []string `json:"hints"`
}

// InterviewSessionQuestionsResponse represents the response for generated questions
type InterviewSessionQuestionsResponse struct {
	SessionID int                  `json:"sessionId"`
	Questions []InterviewQuestion  `json:"questions"`
}

// InterviewSessionDetailsResponse represents detailed session information
type InterviewSessionDetailsResponse struct {
	models.InterviewSession
}
