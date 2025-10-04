# Simple Go GraphQL Backend

A simple Go backend with GraphQL that provides a "hello world" endpoint with mock data.

## Prerequisites

You'll need to install Go first:

### Install Go on macOS:
```bash
# Using Homebrew (recommended)
brew install go

# Or download from https://golang.org/dl/
```

### Verify Installation:
```bash
go version
```

## Setup and Run

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Run the server:**
   ```bash
   go run main.go
   ```

3. **Access the server:**
   - Web interface: http://localhost:8080
   - GraphQL endpoint: http://localhost:8080/graphql

## Testing the GraphQL Endpoint

### Using curl:
```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "{ hello { message status count } }"}'
```

### Expected Response:
```json
{
  "data": {
    "hello": {
      "message": "Hello from GraphQL!",
      "status": "success",
      "count": 42
    }
  }
}
```

## Project Structure

```
stormhacks-BE/
├── go.mod          # Go module dependencies
├── main.go         # Main server code
└── README.md       # This file
```

## Features

- ✅ Simple GraphQL schema with hello world endpoint
- ✅ Mock data (no database required)
- ✅ CORS enabled for web requests
- ✅ Web interface for testing
- ✅ Clean Go module structure
- ✅ Runs on port 8080# stormhacks-BE
