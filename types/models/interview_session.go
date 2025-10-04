package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"stormhacks-be/types/enums"
)

// InterviewSession represents an interview session stored in MongoDB
type InterviewSession struct {
	ID                   primitive.ObjectID        `bson:"_id,omitempty" json:"id"`
	SessionID            int                       `bson:"session_id" json:"sessionId"`
	ParsedResumeText    string                    `bson:"parsed_resume_text" json:"parsedResumeText"`
	JobTitle            string                    `bson:"job_title" json:"jobTitle"`
	JobInfo             string                    `bson:"job_info" json:"jobInfo"`
	CompanyName         *string                   `bson:"company_name,omitempty" json:"companyName,omitempty"`
	AdditionalInfo      *string                   `bson:"additional_info,omitempty" json:"additionalInfo,omitempty"`
	TypeOfInterview     *string                   `bson:"type_of_interview,omitempty" json:"typeOfInterview,omitempty"`
	BehaviouralTopics   []enums.BehaviouralTopic  `bson:"behavioural_topics" json:"behaviouralTopics"`
	TechnicalDifficulty *string                   `bson:"technical_difficulty,omitempty" json:"technicalDifficulty,omitempty"`
	CreatedAt           time.Time                 `bson:"created_at" json:"createdAt"`
}
