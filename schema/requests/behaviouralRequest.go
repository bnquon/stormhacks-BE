package requests

import (
	"github.com/graphql-go/graphql"
)

// BehaviouralFeedbackRequest defines the input type for submitting behavioural feedback
var BehaviouralFeedbackRequest = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "BehaviouralFeedbackRequest",
	Fields: graphql.InputObjectConfigFieldMap{
		"sessionId": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.Int),
			Description: "Session ID to provide feedback for",
		},
		"responses": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.NewList(QuestionResponseRequest)),
			Description: "List of question responses",
		},
	},
})

// QuestionResponseRequest defines the input type for a single question response
var QuestionResponseRequest = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "QuestionResponseRequest",
	Fields: graphql.InputObjectConfigFieldMap{
		"question": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The interview question",
		},
		"response": &graphql.InputObjectFieldConfig{
			Type:        graphql.NewNonNull(graphql.String),
			Description: "The candidate's response to the question",
		},
	},
})
