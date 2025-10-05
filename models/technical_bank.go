package models

import (
	"stormhacks-be/types/enums"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TestCase struct {
	Input          string `bson:"input" json:"input"`
	ExpectedOutput string `bson:"expectedOutput" json:"expectedOutput"`
}

type TechnicalQuestion struct {
	Question     string     `bson:"question" json:"question"`
	Description  string     `bson:"description" json:"description"`
	FunctionName string     `bson:"functionName" json:"functionName"`
	TestCases    []TestCase `bson:"testCases" json:"testCases"`
}

type TechnicalBank struct {
	ID         primitive.ObjectID        `bson:"_id,omitempty" json:"id"`
	Difficulty enums.TechnicalDifficulty `bson:"difficulty" json:"difficulty"`
	Question   TechnicalQuestion         `bson:"question" json:"question"`
}