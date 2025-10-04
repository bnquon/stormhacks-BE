package responses

// QuestionFeedback represents feedback for a single question response
type QuestionFeedback struct {
	Question  string   `json:"question"`
	Response  string   `json:"response"`
	Score     int      `json:"score"`     // Score out of 10
	Feedback  string   `json:"feedback"`  // Detailed feedback
	Suggestions []string `json:"suggestions"` // Improvement suggestions
}

// BehaviouralFeedbackResponse represents the complete feedback response
type BehaviouralFeedbackResponse struct {
	SessionID         int                `json:"sessionId"`
	QuestionFeedback []QuestionFeedback `json:"questionFeedback"`
}
