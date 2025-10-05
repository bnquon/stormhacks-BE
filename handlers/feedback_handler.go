package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"stormhacks-be/types/requests"
)

// FeedbackHandler handles feedback-related HTTP requests
type FeedbackHandler struct {
	interviewService InterviewServiceInterface
}

// NewFeedbackHandler creates a new feedback handler
func NewFeedbackHandler(interviewService InterviewServiceInterface) *FeedbackHandler {
	return &FeedbackHandler{
		interviewService: interviewService,
	}
}

// GenerateFeedback handles POST /api/interview/feedback
func (h *FeedbackHandler) GenerateFeedback(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow POST requests
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse request body
	var input requests.InterviewFeedbackInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if err := h.validateFeedbackInput(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate feedback
	response, err := h.interviewService.GenerateInterviewFeedback(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// validateFeedbackInput validates the input data
func (h *FeedbackHandler) validateFeedbackInput(input requests.InterviewFeedbackInput) error {
	if input.SessionID == "" {
		return errors.New("sessionId is required")
	}
	if len(input.InterviewQuestionsWithAnswers) == 0 {
		return errors.New("interviewQuestionsWithAnswers cannot be empty")
	}
	
	// Validate each question-answer pair
	for _, qa := range input.InterviewQuestionsWithAnswers {
		if qa.Question == "" {
			return errors.New("question cannot be empty")
		}
		if qa.Answer == "" {
			return errors.New("answer cannot be empty")
		}
	}
	
	return nil
}
