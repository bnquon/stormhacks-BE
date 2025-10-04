package responses

type QuestionWithFeedback struct {
	Question string `json:"question"`
	Score int `json:"score"` // 1-10
	Strengths []string `json:"strengths"` // 3 things you did well
	AreasForImprovement []string `json:"areasForImprovement"` // 3 areas to improve
}

type InterviewFeedbackResponse struct {
	SessionID int `json:"sessionId"`
	InterviewQuestionFeedback []QuestionWithFeedback `json:"interviewQuestionFeedback"`
}