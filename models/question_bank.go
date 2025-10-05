package models

import (
	"stormhacks-be/types/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BehavioralQuestionBank represents a behavioral interview question
type QuestionBank struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Question         string             `bson:"question" json:"question"`
	BehavioralTopic  enums.BehaviouralTopic `bson:"behavioralTopic" json:"behavioralTopic"`
}
