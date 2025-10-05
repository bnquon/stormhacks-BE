package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"stormhacks-be/database/migrations"
	"stormhacks-be/database/mongodb"
	"stormhacks-be/handlers"
	"stormhacks-be/repositories"
	"stormhacks-be/services"
)

// ServiceContainer holds all handlers
type ServiceContainer struct {
	InterviewHandler *handlers.InterviewHandler
	FeedbackHandler  *handlers.FeedbackHandler
}

// initializeServices sets up all the service dependencies
func initializeServices() (*ServiceContainer, error) {
	// MongoDB connection
	mongoClient, err := mongodb.NewMongoClient(mongodb.DefaultConfig())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Run database migrations
	if err := migrations.RunMigrations(mongoClient.Database); err != nil {
		log.Printf("Warning: Failed to run migrations: %v", err)
		log.Println("Continuing without migrations...")
	}

	// Create layers
	interviewRepo := repositories.NewInterviewRepository(mongoClient.Database)
	interviewService := services.NewInterviewService(interviewRepo)
	
	// Create handlers
	interviewHandler := handlers.NewInterviewHandler(interviewService)
	feedbackHandler := handlers.NewFeedbackHandler(interviewService)

	return &ServiceContainer{
		InterviewHandler: interviewHandler,
		FeedbackHandler:  feedbackHandler,
	}, nil
}

func main() {
	// Initialize services
	services, err := initializeServices()
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
	http.HandleFunc("/api/interview/session", services.InterviewHandler.CreateInterviewSession)
	http.HandleFunc("/api/interview-questions", services.InterviewHandler.GetInterviewQuestions)
	http.HandleFunc("/api/interview/feedback", services.FeedbackHandler.GenerateFeedback)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Interview API - AI-Powered Interview System</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background-color: #f5f5f5; }
        .container { max-width: 1200px; margin: 0 auto; background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        h1 { color: #2c3e50; border-bottom: 3px solid #3498db; padding-bottom: 10px; }
        h2 { color: #34495e; margin-top: 30px; }
        .feature { background: #ecf0f1; padding: 15px; border-radius: 5px; margin: 10px 0; }
        .ai-badge { background: linear-gradient(45deg, #667eea 0%%, #764ba2 100%%); color: white; padding: 5px 10px; border-radius: 15px; font-size: 12px; font-weight: bold; }
        pre { background: #2c3e50; color: #ecf0f1; padding: 15px; border-radius: 5px; overflow-x: auto; }
        code { background: #ecf0f1; padding: 2px 5px; border-radius: 3px; font-family: 'Courier New', monospace; }
        .endpoint { margin: 10px 0; padding: 10px; border-left: 4px solid #3498db; background: #f8f9fa; }
        .flow { background: #e8f5e8; border: 1px solid #27ae60; border-radius: 5px; padding: 15px; margin: 15px 0; }
        .flow-step { margin: 8px 0; padding-left: 20px; position: relative; }
        .flow-step:before { content: "→"; position: absolute; left: 0; color: #27ae60; font-weight: bold; }
    </style>
</head>
<body>
    <div class="container">
        <h1>AI-Powered Interview API Server</h1>
        <p>Welcome to the StormHacks Interview System! This API provides intelligent interview question generation and feedback using Google Gemini AI.</p>
        
        <div class="feature">
            <h3><span class="ai-badge">AI-POWERED</span> Smart Question Customization</h3>
            <p>The system automatically customizes behavioral interview questions based on the candidate's resume and job description using Google Gemini AI.</p>
        </div>

        <h2>Available API Endpoints:</h2>
        <div class="endpoint">
            <strong><code>GET /health</code></strong> - Health check endpoint
        </div>
        <div class="endpoint">
            <strong><code>POST /api/interview/session</code></strong> - Create interview session with candidate details
        </div>
        <div class="endpoint">
            <strong><code>GET /api/interview-questions?sessionId=123</code></strong> - Get AI-customized interview questions
        </div>
        <div class="endpoint">
            <strong><code>POST /api/interview/feedback</code></strong> - Generate AI-powered interview feedback
        </div>

        <h2>How It Works:</h2>
        <div class="flow">
            <div>Create interview session with resume text and job details</div>
            <div">System fetches base behavioral questions from database</div>
            <div">Google Gemini AI customizes questions based on job role and resume</div>
            <div">Return personalized interview questions</div>
            <div">Generate detailed feedback using AI analysis</div>
        </div>

        <h2>Test the API:</h2>
        
        <h3>1. Create Interview Session:</h3>
        <pre>curl -X POST http://localhost:8080/api/interview/session \\
  -H "Content-Type: application/json" \\
  -d '{
    "sessionId": 123,
    "parsedResumeText": "Experienced software engineer with 5 years in React and Node.js development. Led multiple teams and delivered scalable applications.",
    "jobTitle": "Senior Software Engineer",
    "jobInfo": "We are looking for a senior engineer to lead our frontend team and build scalable React applications.",
    "companyName": "TechCorp",
    "additionalInfo": "Remote work, competitive benefits",
    "typeOfInterview": "behavioral",
    "behaviouralTopics": ["Leadership", "Problem Solving", "Adaptability"],
    "technicalDifficulty": "intermediate"
  }'</pre>

        <h3>2. Get AI-Customized Questions:</h3>
        <pre>curl -X GET "http://localhost:8080/api/interview-questions?sessionId=123"</pre>
        <p><em>Returns questions tailored to the candidate's React/Node.js background and senior engineer role.</em></p>

        <h3>3. Generate Interview Feedback:</h3>
        <pre>curl -X POST http://localhost:8080/api/interview/feedback \\
  -H "Content-Type: application/json" \\
  -d '{
    "sessionId": 123,
    "interviewQuestionsWithAnswers": [
      {
        "question": "Tell me about a time you led a technical team",
        "answer": "I led a team of 5 developers to migrate our legacy system to React..."
      }
    ]
  }'</pre>

        <h2>Features:</h2>
        <ul>
            <li><strong>AI Question Customization:</strong> Questions are tailored based on job requirements and candidate background</li>
            <li><strong>Behavioral Topics:</strong> 9 comprehensive behavioral interview categories</li>
            <li><strong>Smart Feedback:</strong> Detailed AI analysis with scores and improvement suggestions</li>
            <li><strong>MongoDB Integration:</strong> Persistent storage with automatic migrations</li>
            <li><strong>RESTful API:</strong> Clean, well-documented endpoints</li>
        </ul>

        <div class="feature">
            <h3>Ready to Start?</h3>
            <p>Create your first interview session and experience AI-powered question generation!</p>
        </div>
    </div>
</body>
</html>
		`)
	})

	// Start the server
	fmt.Println("🚀 AI-Powered Interview API server starting on http://localhost:8080")
	fmt.Println("🏥 Health check: http://localhost:8080/health")
	fmt.Println("🌐 Web interface: http://localhost:8080")
	fmt.Println("📝 Interview API: http://localhost:8080/api/interview/session")
	fmt.Println("🤖 AI Questions: http://localhost:8080/api/interview-questions")
	fmt.Println("💬 AI Feedback: http://localhost:8080/api/interview/feedback")
	fmt.Println("✨ Powered by Google Gemini AI for intelligent question customization!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
