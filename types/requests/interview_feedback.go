package requests

type QuestionWithAnswer struct {
	Question string `json:"question" validate:"required"`
	Answer string `json:"answer" validate:"required"`
}

type InterviewFeedbackInput struct {
	SessionID int `json:"sessionId" validate:"required"`
	InterviewQuestionsWithAnswers []QuestionWithAnswer `json:"interviewQuestionsWithAnswers" validate:"required"`
}