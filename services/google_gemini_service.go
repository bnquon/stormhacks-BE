package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"stormhacks-be/types/models"
	"stormhacks-be/types/requests"
	"stormhacks-be/types/responses"
	"google.golang.org/genai"
)

// GoogleGeminiService handles Gemini AI interactions
type GoogleGeminiService struct {
	client *genai.Client
}

// NewGoogleGeminiService creates a new Gemini service
func NewGoogleGeminiService() *GoogleGeminiService {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal("Failed to create Gemini client:", err)
	}
	
	return &GoogleGeminiService{
		client: client,
	}
}

// GenerateInterviewFeedback evaluates interview responses using Gemini
func (s *GoogleGeminiService) GenerateInterviewFeedback(session *models.InterviewSession, interviewQuestionsWithAnswers []requests.QuestionWithAnswer) (*responses.InterviewFeedbackResponse, error) {
	ctx := context.Background()
	
	// Build the system prompt
	prompt := s.buildEvaluationPrompt(session, interviewQuestionsWithAnswers)
	
	// Call Gemini API
	result, err := s.client.Models.GenerateContent(
		ctx,
		"gemini-2.0-flash-exp",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content with Gemini: %w", err)
	}
	
	// Parse the response
	feedbackResponse, err := s.parseGeminiResponse(result.Text(), session.SessionID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Gemini response: %w", err)
	}
	
	return feedbackResponse, nil
}

// buildEvaluationPrompt creates the system prompt for Gemini
func (s *GoogleGeminiService) buildEvaluationPrompt(session *models.InterviewSession, questionsWithAnswers []requests.QuestionWithAnswer) string {
	// Build behavioral topics string
	topicsStr := ""
	for i, topic := range session.BehaviouralTopics {
		if i > 0 {
			topicsStr += ", "
		}
		topicsStr += string(topic)
	}
	
	// Build questions and answers string
	qaString := ""
	for i, qa := range questionsWithAnswers {
		qaString += fmt.Sprintf("Question %d: %s\nAnswer %d: %s\n\n", i+1, qa.Question, i+1, qa.Answer)
	}
	
	prompt := fmt.Sprintf(`
You are an expert behavioral interview evaluator. Your task is to evaluate interview responses based on the job context and behavioral topics.

JOB CONTEXT:
- Job Title: %s
- Company: %s
- Job Description: %s
- Resume Summary: %s
- Behavioral Topics Focus: %s

INTERVIEW QUESTIONS AND ANSWERS:
%s

EVALUATION CRITERIA:
1. STAR Method Usage (Situation, Task, Action, Result)
2. Relevance to the question asked
3. Specificity and detail in examples
4. Leadership demonstration
5. Problem-solving approach
6. Communication skills
7. Alignment with job requirements

SCORING SCALE:
- 1-3: Poor (vague, no examples, doesn't answer question)
- 4-5: Below Average (some examples but weak structure)
- 6-7: Good (solid examples, some STAR elements)
- 8-9: Excellent (strong examples, good STAR structure)
- 10: Outstanding (exceptional examples, perfect STAR method)

GENERATE A HIREABILITY SCORE BETWEEN 0 AND 100 BASED ON THE OVERALL INTERVIEW RESPONSES NOT FOR EACH QUESTION.
- 0-20: Not likely to be a good fit
- 21-40: Likely to be a good fit
- 41-60: Very likely to be a good fit
- 61-80: Extremely likely to be a good fit
- 81-100: Perfect fit

RESPONSE FORMAT:
Return ONLY a valid JSON object in this exact format (no markdown, no code blocks, no backticks):
{
  "sessionId": %d,
  "interviewQuestionFeedback": [
    {
      "question": "exact question text",
      "score": 8,
      "strengths": [
        "Specific strength 1 - what they did well",
        "Specific strength 2 - what they did well", 
        "Specific strength 3 - what they did well"
      ],
      "areasForImprovement": [
        "Specific area 1 - what they could improve",
        "Specific area 2 - what they could improve",
        "Specific area 3 - what they could improve"
      ]
    }
  ],
  "hireAbilityScore": 80
}

CRITICAL: The hireAbilityScore field should ONLY appear once at the top level of the response. 
DO NOT include hireAbilityScore inside any individual question objects.
Each question object should ONLY have: question, score, strengths, areasForImprovement.

IMPORTANT:
- Provide specific, actionable feedback
- Reference the job context in your evaluation
- Focus on behavioral competencies
- Be constructive and professional
- For each question, provide exactly 3 strengths and 3 areas for improvement
- Make strengths and improvements specific to the answer given
- Generate ONE overall hireability score (0-100) for the entire interview, not per question
- The hireability score should be at the top level of the response, not inside each question
- Return ONLY the JSON object, no markdown formatting, no code blocks, no backticks
- Do not wrap the JSON in json or any other formatting

FINAL REMINDER: Each question object should contain ONLY these 4 fields:
- question
- score  
- strengths
- areasForImprovement

DO NOT include hireAbilityScore in any question object. It should only appear once at the top level.
`, 
		session.JobTitle,
		session.CompanyName,
		session.JobInfo,
		session.ParsedResumeText,
		topicsStr,
		qaString,
		session.SessionID,
	)
	
	return prompt
}

// parseGeminiResponse parses the JSON response from Gemini
func (s *GoogleGeminiService) parseGeminiResponse(responseText string, sessionID int) (*responses.InterviewFeedbackResponse, error) {
	// Clean the response text - remove markdown formatting
	cleanedText := s.cleanJsonResponse(responseText)
	
	var feedbackResponse responses.InterviewFeedbackResponse
	
	err := json.Unmarshal([]byte(cleanedText), &feedbackResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON response: %w. Response: %s", err, cleanedText)
	}
	
	// Ensure sessionID is set correctly
	feedbackResponse.SessionID = sessionID
	
	return &feedbackResponse, nil
}

// cleanJsonResponse removes markdown formatting from JSON response
func (s *GoogleGeminiService) cleanJsonResponse(responseText string) string {
	// Remove markdown code blocks
	cleaned := responseText
	
	// Remove ```json and ``` markers
	cleaned = strings.ReplaceAll(cleaned, "```json", "")
	cleaned = strings.ReplaceAll(cleaned, "```", "")
	
	// Remove any leading/trailing whitespace
	cleaned = strings.TrimSpace(cleaned)
	
	// Find the first { and last } to extract just the JSON
	start := strings.Index(cleaned, "{")
	end := strings.LastIndex(cleaned, "}")
	
	if start != -1 && end != -1 && end > start {
		cleaned = cleaned[start : end+1]
	}
	
	return cleaned
}