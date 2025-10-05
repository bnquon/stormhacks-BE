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
	http.HandleFunc("/api/technical-question", services.InterviewHandler.GetTechnicalQuestion)
	http.HandleFunc("/api/hint", services.InterviewHandler.GenerateHint)
	http.HandleFunc("/api/execute-code", services.InterviewHandler.ExecuteCode)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>Interview API Documentation</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 40px; background-color: #f5f5f5; }
        .container { max-width: 1200px; margin: 0 auto; background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
        h1 { color: #2c3e50; border-bottom: 3px solid #3498db; padding-bottom: 10px; }
        h2 { color: #34495e; margin-top: 30px; }
        .endpoint { margin: 20px 0; padding: 20px; border: 1px solid #ddd; border-radius: 5px; background: #f8f9fa; }
        .method { display: inline-block; padding: 4px 8px; border-radius: 3px; font-weight: bold; margin-right: 10px; }
        .get { background: #28a745; color: white; }
        .post { background: #007bff; color: white; }
        pre { background: #2c3e50; color: #ecf0f1; padding: 15px; border-radius: 5px; overflow-x: auto; margin: 10px 0; }
        code { background: #ecf0f1; padding: 2px 5px; border-radius: 3px; font-family: 'Courier New', monospace; }
        .url { color: #007bff; font-weight: bold; }
        .response { background: #e8f5e8; border: 1px solid #27ae60; padding: 10px; border-radius: 5px; margin: 10px 0; }
        .payload { background: #fff3cd; border: 1px solid #ffeaa7; padding: 10px; border-radius: 5px; margin: 10px 0; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Interview API Documentation</h1>
        <p>AI-powered interview system with Google Gemini integration for question customization and feedback generation.</p>

        <div class="endpoint">
            <h2><span class="method get">GET</span><span class="url">/health</span></h2>
            <p><strong>Description:</strong> Health check endpoint to verify API status</p>
            <p><strong>Request:</strong></p>
            <pre>curl -X GET http://localhost:8080/health</pre>
            <p><strong>Response:</strong></p>
            <div class="response">
                <pre>{
  "status": "ok",
  "message": "Interview API is running"
}</pre>
            </div>
        </div>

        <div class="endpoint">
            <h2><span class="method post">POST</span><span class="url">/api/interview/session</span></h2>
            <p><strong>Description:</strong> Create a new interview session with candidate details</p>
            <p><strong>Request:</strong></p>
            <pre>curl -X POST http://localhost:8080/api/interview/session \\
  -H "Content-Type: application/json" \\
  -d '{...}'</pre>
            <p><strong>Payload:</strong></p>
            <div class="payload">
                <pre>{
  "parsedResumeText": "Experienced software engineer with 5 years in React and Node.js development. Led multiple teams and delivered scalable applications.",
  "jobTitle": "Senior Software Engineer",
  "jobInfo": "We are looking for a senior engineer to lead our frontend team and build scalable React applications.",
  "companyName": "TechCorp",
  "additionalInfo": "Remote work, competitive benefits",
  "typeOfInterview": "behavioral",
  "behaviouralTopics": ["Leadership", "Problem Solving", "Adaptability"],
  "technicalDifficulty": "Medium"
}</pre>
            </div>
            <p><strong>Response:</strong></p>
            <div class="response">
                <pre>{
  "sessionId": "550e8400-e29b-41d4-a716-446655440000"
}</pre>
            </div>
        </div>

        <div class="endpoint">
            <h2><span class="method get">GET</span><span class="url">/api/interview-questions</span></h2>
            <p><strong>Description:</strong> Get AI-customized interview questions based on session</p>
            <p><strong>Parameters:</strong></p>
            <ul>
                <li><code>sessionId</code> (required) - The session ID created in step 1</li>
            </ul>
            <p><strong>Request:</strong></p>
            <pre>curl -X GET "http://localhost:8080/api/interview-questions?sessionId=550e8400-e29b-41d4-a716-446655440000"</pre>
            <p><strong>Response:</strong></p>
            <div class="response">
                <pre>{
  "sessionId": "550e8400-e29b-41d4-a716-446655440000",
  "questions": [
    {
      "id": "q1",
      "topic": "Leadership",
      "question": "TechCorp values mentorship. Tell me about a time you successfully guided a junior React developer to overcome a performance bottleneck in a complex component. How did you approach their learning, and what was the final impact on the application's performance?",
      "hints": [
        "What was the specific performance issue the junior developer was facing?",
        "What specific mentoring techniques or approaches did you use?",
        "What measurable improvements were achieved in the component's performance?"
      ]
    },
    {
      "id": "q2",
      "topic": "Problem Solving",
      "question": "TechCorp prioritizes delivering high-performing applications. Tell me about a time you used performance monitoring tools and data analysis to identify and resolve a significant performance bottleneck in a React application.",
      "hints": [
        "What specific metrics did you track?",
        "What was your diagnostic process?",
        "What impact did your solution have on the user experience?"
      ]
    },
    {
      "id": "q3",
      "topic": "Adaptability",
      "question": "At TechCorp, you might be communicating technical details to different audiences. Describe how you adapt your communication style when explaining complex React frontend architectures.",
      "hints": [
        "What specific adjustments did you make for each audience?",
        "How did you ensure everyone was aligned and understood the information?",
        "What was the outcome of your communication approach?"
      ]
    }
  ]
}</pre>
            </div>
        </div>

        <div class="endpoint">
            <h2><span class="method post">POST</span><span class="url">/api/interview/feedback</span></h2>
            <p><strong>Description:</strong> Generate AI-powered interview feedback based on candidate responses</p>
            <p><strong>Request:</strong></p>
            <pre>curl -X POST http://localhost:8080/api/interview/feedback \\
  -H "Content-Type: application/json" \\
  -d '{...}'</pre>
            <p><strong>Payload:</strong></p>
            <div class="payload">
                <pre>{
  "sessionId": "550e8400-e29b-41d4-a716-446655440000",
  "interviewQuestionsWithAnswers": [
    {
      "question": "Tell me about a time you led a technical team",
      "answer": "I led a team of 5 developers to migrate our legacy system to React. We faced performance issues and tight deadlines. I organized daily standups, delegated specific tasks based on each developer's strengths, and implemented code review processes. The result was a 40% improvement in application performance and the team gained valuable React expertise."
    },
    {
      "question": "Describe a time you solved a complex technical problem",
      "answer": "Our React application was experiencing memory leaks causing crashes. I used React DevTools profiler to identify the issue was in our useEffect cleanup. I implemented proper dependency arrays and cleanup functions, reducing memory usage by 60% and eliminating crashes."
    }
  ]
}</pre>
            </div>
            <p><strong>Response:</strong></p>
            <div class="response">
                <pre>{
  "sessionId": "550e8400-e29b-41d4-a716-446655440000",
  "interviewQuestionFeedback": [
    {
      "question": "Tell me about a time you led a technical team",
      "score": 8,
      "strengths": [
        "Clearly articulated the situation and task at hand",
        "Demonstrated leadership by organizing standups and delegating tasks",
        "Quantified results with specific improvements"
      ],
      "areasForImprovement": [
        "Could elaborate on the specific technical challenges faced",
        "Could provide more detail about the team's individual contributions",
        "Could discuss lessons learned from the experience"
      ]
    },
    {
      "question": "Describe a time you solved a complex technical problem",
      "score": 9,
      "strengths": [
        "Showed strong technical problem-solving skills",
        "Used appropriate tools for diagnosis",
        "Provided specific, measurable results"
      ],
      "areasForImprovement": [
        "Could explain the broader impact on the application",
        "Could discuss preventive measures implemented"
      ]
    }
  ],
  "hireAbilityScore": 87,
  "overallFeedback": [
    "Strong technical leadership with clear communication and measurable results",
    "Excellent problem-solving approach using appropriate tools and methodologies",
    "Consider expanding on broader impact and preventive measures for even stronger responses"
  ]
}</pre>
            </div>
        </div>

        <div class="endpoint">
            <h2><span class="method get">GET</span><span class="url">/api/technical-question</span></h2>
            <p><strong>Description:</strong> Get a random technical question by difficulty level</p>
            <p><strong>Parameters:</strong></p>
            <ul>
                <li><code>difficulty</code> (required) - The difficulty level: Easy, Medium, or Hard</li>
            </ul>
            <p><strong>Request:</strong></p>
            <pre>curl -X GET "http://localhost:8080/api/technical-question?difficulty=Medium"</pre>
            <p><strong>Response:</strong></p>
            <div class="response">
                <pre>{
  "id": "507f1f77bcf86cd799439011",
  "difficulty": "Medium",
  "question": {
    "questionId": "68e205dadb8a0fc4ec6924e9",
    "description": "Given two strings, find the length of the longest common subsequence. A subsequence is a sequence that appears in the same relative order, but not necessarily contiguous.",
    "functionName": "longestCommonSubsequence",
    "testCases": [
      {
        "input": "ABCDGH, AEDFHR",
        "expectedOutput": "3 (ADH)"
      },
      {
        "input": "AGGTAB, GXTXAYB",
        "expectedOutput": "4 (GTAB)"
      }
    ]
  }
}</pre>
            </div>
        </div>

        <div class="endpoint">
            <h2><span class="method post">POST</span><span class="url">/api/hint</span></h2>
            <p><strong>Description:</strong> Generate AI-powered hints for interview responses with text-to-speech support</p>
            <p><strong>Request:</strong></p>
            <pre>curl -X POST http://localhost:8080/api/hint \\
  -H "Content-Type: application/json" \\
  -d '{...}'</pre>
            <p><strong>Payload:</strong></p>
            <div class="payload">
                <pre>{
  "sessionId": "550e8400-e29b-41d4-a716-446655440000",
  "questionId": "68e205dadb8a0fc4ec6924e9",
  "previousHints": ["Think about dynamic programming", "Consider the recursive relationship"],
  "userCode": "function lcs(str1, str2) {\n  // My current attempt\n  return '';\n}",
  "userSpeech": "I'm trying to implement this function but I'm not sure where to start. Can you help me?"
}</pre>
            </div>
            <p><strong>Response:</strong></p>
            <div class="response">
                <pre>{
  "sessionId": "550e8400-e29b-41d4-a716-446655440000",
  "conversationalHint": "Great start! I can see you're thinking about this problem. Let me guide you - consider what happens when you compare characters at each position. What would you do if the characters match versus when they don't match?",
  "hintSummary": "Consider character comparison logic for matching vs non-matching cases"
}</pre>
            </div>
        </div>

        <div class="endpoint">
            <h2><span class="method post">POST</span><span class="url">/api/hint</span></h2>
            <p><strong>Description:</strong> Generate AI-powered hints for interview responses with text-to-speech support</p>
            <p><strong>Request:</strong></p>
            <pre>curl -X POST http://localhost:8080/api/hint \\
  -H "Content-Type: application/json" \\
  -d '{...}'</pre>
            <p><strong>Payload:</strong></p>
            <div class="payload">
                <pre>{
  "sessionId": "550e8400-e29b-41d4-a716-446655440000",
  "questionId": "68e205dadb8a0fc4ec6924e9",
  "previousHints": ["Think about dynamic programming", "Consider the recursive relationship"],
  "userCode": "function lcs(str1, str2) {\n  // My current attempt\n  return '';\n}",
  "userSpeech": "I'm trying to implement this function but I'm not sure where to start. Can you help me?"
}</pre>
            </div>
            <p><strong>Response:</strong></p>
            <div class="response">
                <pre>{
  "sessionId": "550e8400-e29b-41d4-a716-446655440000",
  "conversationalHint": "Great start! I can see you're thinking about this problem. Let me guide you - consider what happens when you compare characters at each position. What would you do if the characters match versus when they don't match?",
  "hintSummary": "Consider character comparison logic for matching vs non-matching cases"
}</pre>
            </div>
        </div>

        <div class="endpoint">
            <h2><span class="method post">POST</span><span class="url">/api/execute-code</span></h2>
            <p><strong>Description:</strong> Execute user-submitted code against test cases for technical questions</p>
            <p><strong>Request:</strong></p>
            <pre>curl -X POST http://localhost:8080/api/execute-code \\
  -H "Content-Type: application/json" \\
  -d '{...}'</pre>
            <p><strong>Payload:</strong></p>
            <div class="payload">
                <pre>{
  "questionId": "68e205dadb8a0fc4ec6924e9",
  "code": "def rottenOranges(grid):\n    if not grid or not grid[0]:\n        return -1\n    \n    m, n = len(grid), len(grid[0])\n    queue = []\n    fresh = 0\n    \n    # Find all rotten oranges and count fresh ones\n    for i in range(m):\n        for j in range(n):\n            if grid[i][j] == 2:\n                queue.append((i, j, 0))\n            elif grid[i][j] == 1:\n                fresh += 1\n    \n    if fresh == 0:\n        return 0\n    \n    directions = [(0,1), (1,0), (0,-1), (-1,0)]\n    \n    while queue:\n        i, j, time = queue.pop(0)\n        \n        for di, dj in directions:\n            ni, nj = i + di, j + dj\n            if 0 <= ni < m and 0 <= nj < n and grid[ni][nj] == 1:\n                grid[ni][nj] = 2\n                fresh -= 1\n                if fresh == 0:\n                    return time + 1\n                queue.append((ni, nj, time + 1))\n    \n    return -1",
  "language": "python"
}</pre>
            </div>
            <p><strong>Response (Success):</strong></p>
            <div class="response">
                <pre>{
  "questionId": "68e205dadb8a0fc4ec6924e9",
  "code": "def rottenOranges(grid):...",
  "language": "python",
  "output": "4\n-1\n0",
  "error": "",
  "executionTime": 200,
  "success": true
}</pre>
            </div>
            <p><strong>Response (Compilation Error):</strong></p>
            <div class="response">
                <pre>{
  "questionId": "68e205dadb8a0fc4ec6924e9",
  "code": "def rottenOranges(grid):...",
  "language": "python",
  "output": "",
  "error": "Compilation Error: SyntaxError: expected ':'",
  "executionTime": 50,
  "success": false
}</pre>
            </div>
            <p><strong>Response (Wrong Answer):</strong></p>
            <div class="response">
                <pre>{
  "questionId": "68e205dadb8a0fc4ec6924e9",
  "code": "def rottenOranges(grid):...",
  "language": "python",
  "output": "1\n1\n0",
  "error": "Test case 1: Expected '4', got '1'\nTest case 2: Expected '-1', got '1'",
  "executionTime": 150,
  "success": false
}</pre>
            </div>
            <p><strong>Supported Languages:</strong></p>
            <ul>
                <li><code>python</code> - Python 3</li>
                <li><code>javascript</code> - JavaScript (Node.js)</li>
            </ul>
            <p><strong>Error Types:</strong></p>
            <ul>
                <li><strong>Compilation Error</strong> - Syntax errors, missing imports, etc.</li>
                <li><strong>Runtime Error</strong> - Index out of bounds, null pointer, etc.</li>
                <li><strong>Wrong Answer</strong> - Code runs but produces incorrect output</li>
                <li><strong>Execution Error</strong> - Timeout, memory issues, etc.</li>
            </ul>
        </div>


        <div class="endpoint">
            <h2>Available Behavioral Topics</h2>
            <ul>
                <li>General</li>
                <li>Workplace Behavior</li>
                <li>Leadership</li>
                <li>Problem Solving</li>
                <li>Conflict Resolution</li>
                <li>Adaptability</li>
                <li>Time Management</li>
                <li>Customer Focus</li>
                <li>Innovation & Creativity</li>
            </ul>
        </div>

        <div class="endpoint">
            <h2>Interview Types</h2>
            <ul>
                <li>technical</li>
                <li>behavioral</li>
                <li>both</li>
            </ul>
        </div>

        <div class="endpoint">
            <h2>Technical Difficulty Levels</h2>
            <ul>
                <li>Easy</li>
                <li>Medium</li>
                <li>Hard</li>
            </ul>
        </div>
    </div>
</body>
</html>
		`)
	})

	// Start the server
	fmt.Println("AI-Powered Interview API server starting on http://localhost:8080")
	fmt.Println("Health check: http://localhost:8080/health")
	fmt.Println("Web interface: http://localhost:8080")
	fmt.Println("Interview API: http://localhost:8080/api/interview/session")
	fmt.Println("AI Questions: http://localhost:8080/api/interview-questions")
	fmt.Println("AI Feedback: http://localhost:8080/api/interview/feedback")
	fmt.Println("Technical Questions: http://localhost:8080/api/technical-question")
	fmt.Println("Hint Generation: http://localhost:8080/api/hint")
	fmt.Println("Code Execution: http://localhost:8080/api/execute-code")
	fmt.Println("Powered by Google Gemini AI for intelligent question customization and hints!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
