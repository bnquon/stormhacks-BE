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

// GetInterviewQuestions handles GET /api/interview-questions
func (h *InterviewHandler) GetInterviewQuestions(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow GET requests
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get sessionId from query parameters
	sessionIdStr := r.URL.Query().Get("sessionId")
	if sessionIdStr == "" {
		http.Error(w, "sessionId query parameter is required", http.StatusBadRequest)
		return
	}

	// Use sessionId as string (UUID)
	sessionId := sessionIdStr

	// Get interview questions
	response, err := h.interviewService.GenerateInterviewQuestions(sessionId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// validateInterviewSessionInput validates the input data
func (h *InterviewHandler) validateInterviewSessionInput(input requests.InterviewSessionInput) error {
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

// GetTechnicalQuestion handles GET /api/technical-question
func (h *InterviewHandler) GetTechnicalQuestion(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow GET requests
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get difficulty from query parameters
	difficulty := r.URL.Query().Get("difficulty")
	if difficulty == "" {
		http.Error(w, "difficulty query parameter is required", http.StatusBadRequest)
		return
	}

	// Get technical question
	question, err := h.interviewService.GetTechnicalQuestion(difficulty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(question)
}

// ExecuteCode handles POST /api/execute-code
func (h *InterviewHandler) ExecuteCode(w http.ResponseWriter, r *http.Request) {
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
	var input requests.ExecuteTechnicalInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if err := h.validateExecuteTechnicalInput(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execute code
	response, err := h.interviewService.ExecuteCode(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GenerateHint handles POST /api/hint
func (h *InterviewHandler) GenerateHint(w http.ResponseWriter, r *http.Request) {
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
	var input requests.HintRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate input
	if err := h.validateHintRequest(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate hints
	response, err := h.interviewService.GenerateHint(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// validateExecuteTechnicalInput validates the input data
func (h *InterviewHandler) validateExecuteTechnicalInput(input requests.ExecuteTechnicalInput) error {
	if input.QuestionID == "" {
		return errors.New("questionId is required")
	}
	if input.Code == "" {
		return errors.New("code is required")
	}
	return nil
}

// validateHintRequest validates the hint request input
func (h *InterviewHandler) validateHintRequest(input requests.HintRequest) error {
	if input.SessionID == "" {
		return errors.New("sessionId is required")
	}
	if input.Question == "" {
		return errors.New("question is required")
	}
	if input.UserCode == "" {
		return errors.New("userCode is required")
	}
	if input.UserSpeech == "" {
		return errors.New("userSpeech is required")
	}
	return nil
}

