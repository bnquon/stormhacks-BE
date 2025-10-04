package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stormhacks-be/database/mongodb"
	"stormhacks-be/handlers"
	"stormhacks-be/repositories"
	"stormhacks-be/services"
)

// initializeServices sets up all the service dependencies
func initializeServices() (*handlers.InterviewHandler, error) {
	// MongoDB connection
	mongoClient, err := mongodb.NewMongoClient(mongodb.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Get collection
	collection := mongoClient.GetCollection("interview_sessions")

	// Create layers
	interviewRepo := repositories.NewInterviewRepository(collection)
	interviewService := services.NewInterviewService(interviewRepo)
	interviewHandler := handlers.NewInterviewHandler(interviewService)

	return interviewHandler, nil
}

func main() {
	// Initialize services
	interviewHandler, err := initializeServices()
	if err != nil {
		log.Fatal("Failed to initialize services:", err)
	}

	// Simple health check handler
	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "Interview API is running",
		})
	}

	// Set up routes
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/interview/session", interviewHandler.CreateInterviewSession)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Interview API</title>
</head>
<body>
    <h1>Interview API Server Running!</h1>
    <p>API endpoints available:</p>
    <ul>
        <li><code>GET /health</code> - Health check</li>
        <li><code>POST /api/interview/session</code> - Create interview session</li>
    </ul>
    <h2>Test Interview Session:</h2>
    <pre>curl -X POST http://localhost:8080/api/interview/session \\
  -H "Content-Type: application/json" \\
  -d '{"sessionId": 123, "parsedResumeText": "resume text", "jobTitle": "Software Engineer", "jobInfo": "job description"}'</pre>
</body>
</html>
		`)
	})

	// Start the server
	fmt.Println("🚀 Interview API server starting on http://localhost:8080")
	fmt.Println("🏥 Health check: http://localhost:8080/health")
	fmt.Println("🌐 Web interface: http://localhost:8080")
	fmt.Println("📝 Interview API: http://localhost:8080/api/interview/session")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}