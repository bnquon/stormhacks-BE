package schema

import (
	"fmt"
	"stormhacks-be/schema/requests"
	"stormhacks-be/schema/responses"

	"github.com/graphql-go/graphql"
)

// CreateSchema creates and returns the complete GraphQL schema with the provided resolvers
func CreateSchema(interviewResolvers interface{}) (*graphql.Schema, error) {
	// Type assertion to get the resolvers
	resolvers, ok := interviewResolvers.(interface {
		GetInterviewSession(graphql.ResolveParams) (interface{}, error)
		GenerateInterviewQuestions(graphql.ResolveParams) (interface{}, error)
		GetBehaviouralFeedback(graphql.ResolveParams) (interface{}, error)
		CreateInterviewSession(graphql.ResolveParams) (interface{}, error)
		SubmitBehaviouralFeedback(graphql.ResolveParams) (interface{}, error)
	})
	
	if !ok {
		return nil, fmt.Errorf("Invalid resolver type")
	}

	// Define the root query
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"getInterviewSession": &graphql.Field{
				Type:        responses.InterviewSessionResponse,
				Description: "Get an interview session by session ID",
				Args: graphql.FieldConfigArgument{
					"sessionId": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "Session ID to retrieve",
					},
				},
				Resolve: resolvers.GetInterviewSession,
			},
			"generateInterviewQuestions": &graphql.Field{
				Type:        responses.InterviewSessionQuestionsResponse,
				Description: "Generate interview questions for a session",
				Args: graphql.FieldConfigArgument{
					"sessionId": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "Session ID to generate questions for",
					},
				},
				Resolve: resolvers.GenerateInterviewQuestions,
			},
			"getBehaviouralFeedback": &graphql.Field{
				Type:        responses.BehaviouralFeedbackResponse,
				Description: "Get behavioural feedback for a session",
				Args: graphql.FieldConfigArgument{
					"sessionId": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "Session ID to get feedback for",
					},
				},
				Resolve: resolvers.GetBehaviouralFeedback,
			},
		},
	})

	// Define the root mutation
	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createInterviewSession": &graphql.Field{
				Type:        responses.InterviewSessionResponse,
				Description: "Create a new interview session",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(requests.InterviewSessionRequest),
						Description: "Interview session input data",
					},
				},
				Resolve: resolvers.CreateInterviewSession,
			},
			"submitBehaviouralFeedback": &graphql.Field{
				Type:        responses.BehaviouralFeedbackResponse,
				Description: "Submit behavioural feedback for interview responses",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(requests.BehaviouralFeedbackRequest),
						Description: "Behavioural feedback input data",
					},
				},
				Resolve: resolvers.SubmitBehaviouralFeedback,
			},
		},
	})

	// Create the schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})

	return &schema, err
}
