package responses

import (
	"stormhacks-be/models"
	"stormhacks-be/schema/enums"

	"github.com/graphql-go/graphql"
)

// InterviewSessionResponse defines the output type for interview session data
var InterviewSessionResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "InterviewSession",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "Unique identifier for the interview session",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if session, ok := p.Source.(models.InterviewSession); ok {
					return session.ID.Hex(), nil
				}
				return nil, nil
			},
		},
		"sessionId": &graphql.Field{
			Type:        graphql.Int,
			Description: "Session ID number",
		},
		"parsedResumeText": &graphql.Field{
			Type:        graphql.String,
			Description: "Parsed resume text content",
		},
		"jobTitle": &graphql.Field{
			Type:        graphql.String,
			Description: "Title of the job position",
		},
		"jobInfo": &graphql.Field{
			Type:        graphql.String,
			Description: "Detailed job information",
		},
		"companyName": &graphql.Field{
			Type:        graphql.String,
			Description: "Name of the company",
		},
		"additionalInfo": &graphql.Field{
			Type:        graphql.String,
			Description: "Additional information about the interview",
		},
		"typeOfInterview": &graphql.Field{
			Type:        graphql.String,
			Description: "Type of interview",
		},
		"behaviouralTopics": &graphql.Field{
			Type:        graphql.NewList(enums.BehaviouralTopicEnum),
			Description: "List of behavioural topics",
		},
		"technicalDifficulty": &graphql.Field{
			Type:        graphql.String,
			Description: "Technical difficulty level",
		},
		"createdAt": &graphql.Field{
			Type:        graphql.String,
			Description: "When the session was created",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				if session, ok := p.Source.(models.InterviewSession); ok {
					return session.CreatedAt.Format("2006-01-02T15:04:05Z07:00"), nil
				}
				return nil, nil
			},
		},
	},
})
