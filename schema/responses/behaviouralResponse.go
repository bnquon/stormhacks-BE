package responses

import (
	"github.com/graphql-go/graphql"
)

// QuestionResponse defines the output type for interview questions
var QuestionResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "Question",
	Fields: graphql.Fields{
		"question": &graphql.Field{
			Type:        graphql.String,
			Description: "The interview question",
		},
		"hints": &graphql.Field{
			Type:        graphql.NewList(graphql.String),
			Description: "Hints for answering the question",
		},
	},
})

// InterviewSessionQuestionsResponse defines the output type for interview questions response
var InterviewSessionQuestionsResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "InterviewSessionQuestionsResponse",
	Fields: graphql.Fields{
		"sessionId": &graphql.Field{
			Type:        graphql.Int,
			Description: "Session ID",
		},
		"questions": &graphql.Field{
			Type:        graphql.NewList(QuestionResponse),
			Description: "Generated interview questions",
		},
	},
})

// QuestionFeedbackResponse defines the output type for question feedback
var QuestionFeedbackResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "QuestionFeedback",
	Fields: graphql.Fields{
		"question": &graphql.Field{
			Type:        graphql.String,
			Description: "The interview question",
		},
		"response": &graphql.Field{
			Type:        graphql.String,
			Description: "The candidate's response",
		},
		"score": &graphql.Field{
			Type:        graphql.Int,
			Description: "Score out of 10",
		},
		"feedback": &graphql.Field{
			Type:        graphql.String,
			Description: "Detailed feedback",
		},
		"suggestions": &graphql.Field{
			Type:        graphql.NewList(graphql.String),
			Description: "Improvement suggestions",
		},
	},
})

// BehaviouralFeedbackResponse defines the output type for behavioural feedback
var BehaviouralFeedbackResponse = graphql.NewObject(graphql.ObjectConfig{
	Name: "BehaviouralFeedbackResponse",
	Fields: graphql.Fields{
		"sessionId": &graphql.Field{
			Type:        graphql.Int,
			Description: "Session ID",
		},
		"questionFeedback": &graphql.Field{
			Type:        graphql.NewList(QuestionFeedbackResponse),
			Description: "Feedback for each question",
		},
	},
})
