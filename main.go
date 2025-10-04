package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"stormhacks-be/database/mongodb"
	"stormhacks-be/resolvers"
	"stormhacks-be/schema"

	"github.com/graphql-go/graphql"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database connection
	if err := mongodb.InitDatabase(); err != nil {
		log.Printf("Warning: Failed to connect to MongoDB: %v", err)
		log.Println("Continuing without database connection...")
	}
	defer mongodb.CloseDatabase()

	// Initialize resolvers
	interviewResolvers := resolvers.NewInterviewResolvers()

	// Get the GraphQL schema
	graphqlSchema, err := schema.CreateSchema(interviewResolvers)
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	// GraphQL handler function
	graphqlHandler := func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Only allow POST requests for GraphQL
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Parse the request body
		var requestBody struct {
			Query string `json:"query"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Execute the GraphQL query
		result := graphql.Do(graphql.Params{
			Schema:        *graphqlSchema,
			RequestString: requestBody.Query,
		})

		// Set content type and return the result
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}

	// Set up routes
	http.HandleFunc("/graphql", graphqlHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
    <title>GraphQL Server</title>
</head>
<body>
    <h1>GraphQL Server Running!</h1>
    <p>Send POST requests to <code>/graphql</code> endpoint</p>
    <h2>Example Query:</h2>
    <pre>{
  "query": "{ getInterviewSession(sessionId: 1) { sessionId jobTitle companyName } }"
}</pre>
    <h2>Example Mutation:</h2>
    <pre>{
  "query": "mutation { createInterviewSession(input: { sessionId: 1, parsedResumeText: \"Sample resume\", jobTitle: \"Software Engineer\", jobInfo: \"Full stack development\" }) { sessionId jobTitle } }"
}</pre>
    <h2>Test with curl:</h2>
    <pre>curl -X POST http://localhost:8080/graphql \\
  -H "Content-Type: application/json" \\
  -d '{"query": "{ getInterviewSession(sessionId: 1) { sessionId jobTitle } }"}'</pre>
</body>
</html>
		`)
	})

	// Get port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	fmt.Printf("🚀 GraphQL server starting on http://localhost:%s\n", port)
	fmt.Printf("📊 GraphQL endpoint: http://localhost:%s/graphql\n", port)
	fmt.Printf("🌐 Web interface: http://localhost:%s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
