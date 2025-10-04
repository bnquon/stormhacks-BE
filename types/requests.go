package types

// InterviewSessionInput represents the input for creating an interview session
type InterviewSessionInput struct {
	SessionID           int      `json:"sessionId" validate:"required"`
	ParsedResumeText    string   `json:"parsedResumeText" validate:"required"`
	JobTitle            string   `json:"jobTitle" validate:"required"`
	JobInfo             string   `json:"jobInfo" validate:"required"`
	CompanyName         *string  `json:"companyName,omitempty"`
	AdditionalInfo      *string  `json:"additionalInfo,omitempty"`
	TypeOfInterview     *string  `json:"typeOfInterview,omitempty"`     // Future implementation
	BehaviouralTopics   []string `json:"behaviouralTopics,omitempty"`   // Default: ["general"]
	TechnicalDifficulty *string  `json:"technicalDifficulty,omitempty"` // Future implementation
}
