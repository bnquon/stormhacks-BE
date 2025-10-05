package requests

import "stormhacks-be/types/enums"

// InterviewSessionInput represents the input for creating an interview session
type InterviewSessionInput struct {
	ParsedResumeText    string                   `json:"parsedResumeText" validate:"required"`
	JobTitle            string                   `json:"jobTitle" validate:"required"`
	JobInfo             string                   `json:"jobInfo" validate:"required"`
	CompanyName         *string                  `json:"companyName,omitempty"`
	AdditionalInfo      *string                  `json:"additionalInfo,omitempty"`
	TypeOfInterview     *string                  `json:"typeOfInterview,omitempty"` // Future implementation
	BehaviouralTopics   []enums.BehaviouralTopic `json:"behaviouralTopics,omitempty"` // Default: ["General"]
	TechnicalDifficulty *enums.TechnicalDifficulty `json:"technicalDifficulty,omitempty"` // Future implementation
}
