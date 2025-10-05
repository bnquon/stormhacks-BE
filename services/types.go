package services

import "stormhacks-be/types/enums"

// CustomizedQuestion represents a question with AI-generated hints
type CustomizedQuestion struct {
	ID              string
	Question        string
	BehavioralTopic enums.BehaviouralTopic
	Hints           []string
}
