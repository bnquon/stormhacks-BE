package requests

import (
	"stormhacks-be/schema/enums"

	"github.com/graphql-go/graphql"
)

// InterviewSessionRequest defines the input type for creating an interview session
var InterviewSessionRequest = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "InterviewSessionRequest",
	Fields: graphql.InputObjectConfigFieldMap{
		"sessionId": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Unique session identifier",
		},
		"parsedResumeText": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Parsed resume text content",
		},
		"jobTitle": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Title of the job position",
		},
		"jobInfo": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "Detailed job information",
		},
		"companyName": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "Name of the company (optional)",
		},
		"additionalInfo": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "Additional information about the interview (optional)",
		},
		"typeOfInterview": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "Type of interview (future implementation)",
		},
		"behaviouralTopics": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewList(enums.BehaviouralTopicEnum),
			Description: "List of behavioural topics to cover",
		},
		"technicalDifficulty": &graphql.InputObjectFieldConfig{
			Type:        graphql.String,
			Description: "Technical difficulty level (future implementation)",
		},
	},
})
