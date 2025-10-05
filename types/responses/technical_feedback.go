package responses

type TechnicalFeedbackResponse struct {
	SessionID string `json:"sessionId"`
	HireAbilityScore int `json:"hireAbilityScore"` // 0-100
	Suggestions []string `json:"suggestions"` // 3 suggestions for improvement
	Strengths []string `json:"strengths"` // 3 things you did well
}