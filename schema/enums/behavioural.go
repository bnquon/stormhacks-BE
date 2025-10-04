package enums

import (
	"github.com/graphql-go/graphql"
)

// BehaviouralTopicEnum defines the available behavioural interview topics
var BehaviouralTopicEnum = graphql.NewEnum(graphql.EnumConfig{
	Name: "BehaviouralTopic",
	Values: graphql.EnumValueConfigMap{
		"GENERAL": &graphql.EnumValueConfig{
			Value:       "general",
			Description: "General behavioural questions",
		},
		"WORKPLACE_BEHAVIOR": &graphql.EnumValueConfig{
			Value:       "workplace_behavior",
			Description: "Workplace behavior and professionalism",
		},
		"LEADERSHIP": &graphql.EnumValueConfig{
			Value:       "leadership",
			Description: "Leadership and management skills",
		},
		"PROBLEM_SOLVING": &graphql.EnumValueConfig{
			Value:       "problem_solving",
			Description: "Problem solving and analytical thinking",
		},
		"CONFLICT_RESOLUTION": &graphql.EnumValueConfig{
			Value:       "conflict_resolution",
			Description: "Conflict resolution and communication",
		},
		"ADAPTABILITY": &graphql.EnumValueConfig{
			Value:       "adaptability",
			Description: "Adaptability and flexibility",
		},
		"TIME_MANAGEMENT": &graphql.EnumValueConfig{
			Value:       "time_management",
			Description: "Time management and organization",
		},
		"CUSTOMER_FOCUS": &graphql.EnumValueConfig{
			Value:       "customer_focus",
			Description: "Customer focus and service orientation",
		},
		"INNOVATION_CREATIVITY": &graphql.EnumValueConfig{
			Value:       "innovation_creativity",
			Description: "Innovation and creativity",
		},
	},
})
