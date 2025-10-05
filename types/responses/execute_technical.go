package responses

// ExecuteTechnicalResponse represents the response for code execution
type ExecuteTechnicalResponse struct {
	QuestionID   string `json:"questionId"`
	Code         string `json:"code"`
	Language     string `json:"language"`
	Output       string `json:"output"`
	Error        string `json:"error,omitempty"`
	ExecutionTime int64 `json:"executionTime"` // in milliseconds
	Success      bool   `json:"success"`
}
