package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// HelloResponse represents our mock data structure
type HelloResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Count   int    `json:"count"`
}

// Mock data
var mockData = HelloResponse{
	Message: "Hello from GraphQL!",
	Status:  "success",
	Count:   42,
}

// Root query resolver
func rootQueryResolver(p graphql.ResolveParams) (interface{}, error) {
	return mockData, nil
}

func main() {
	// Define the GraphQL schema
	helloType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Hello",
		Fields: graphql.Fields{
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"status": &graphql.Field{
				Type: graphql.String,
			},
			"count": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

	// Define the root query
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type:        helloType,
				Description: "Get a hello world message",
				Resolve:     rootQueryResolver,
			},
		},
	})

	// Create the schema
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
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
			Schema:        schema,
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
  "query": "{ hello { message status } }"
}</pre>
    <h2>Test with curl:</h2>
    <pre>curl -X POST http://localhost:8080/graphql \\
  -H "Content-Type: application/json" \\
  -d '{"query": "{ hello { message status } }"}'</pre>
</body>
</html>
		`)
	})

	// Start the server
	fmt.Println("🚀 GraphQL server starting on http://localhost:8080")
	fmt.Println("📊 GraphQL endpoint: http://localhost:8080/graphql")
	fmt.Println("🌐 Web interface: http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
