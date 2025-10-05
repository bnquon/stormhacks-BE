package responses

type QuestionWithFeedback struct {
	Question string `json:"question"`
	Score int `json:"score"` // 1-10
	Strengths []string `json:"strengths"` // 3 things you did well
	AreasForImprovement []string `json:"areasForImprovement"` // 3 areas to improve
	HireAbilityScore int `json:"hireAbilityScore"` // 0-100
}

type InterviewFeedbackResponse struct {
	SessionID string `json:"sessionId"`
	InterviewQuestionFeedback []QuestionWithFeedback `json:"interviewQuestionFeedback"`
	HireAbilityScore int `json:"hireAbilityScore"` // 0-100
}