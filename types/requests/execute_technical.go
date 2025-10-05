package requests

import "stormhacks-be/types/enums"

// ExecuteTechnicalInput represents the input for executing technical code
type ExecuteTechnicalInput struct {
	QuestionID string                `json:"questionId" validate:"required"`
	Code       string                `json:"code" validate:"required"`
	Language   enums.CodingLanguage  `json:"language" validate:"required"`
}
