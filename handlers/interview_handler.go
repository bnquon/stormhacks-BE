package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"stormhacks-be/types/requests"
)


// InterviewHandler handles interview-related HTTP requests
type InterviewHandler struct {
	interviewService InterviewServiceInterface
}

// NewInterviewHandler creates a new interview handler
func NewInterviewHandler(interviewService InterviewServiceInterface) *InterviewHandler {
	return &InterviewHandler{
		interviewService: interviewService,
	}
}

// CreateInterviewSession handles POST /api/interview/session
func (h *InterviewHandler) CreateInterviewSession(w http.ResponseWriter, r *http.Request) {
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
	var input requests.InterviewSessionInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if err := h.validateInterviewSessionInput(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create interview session
	response, err := h.interviewService.CreateInterviewSession(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// validateInterviewSessionInput validates the input data
func (h *InterviewHandler) validateInterviewSessionInput(input requests.InterviewSessionInput) error {
	// Basic validation - you can add more sophisticated validation here
	if input.SessionID <= 0 {
		return errors.New("sessionId must be greater than 0")
	}
	if input.ParsedResumeText == "" {
		return errors.New("parsedResumeText is required")
	}
	if input.JobTitle == "" {
		return errors.New("jobTitle is required")
	}
	if input.JobInfo == "" {
		return errors.New("jobInfo is required")
	}
	return nil
}
