# AI-Powered Interview API

A comprehensive interview system with Google Gemini AI integration for question customization, code execution, and intelligent feedback generation.

## Features

- **AI-Powered Questions**: Customized interview questions based on job context and candidate profile
- **Code Execution**: Execute and validate code submissions against test cases
- **Intelligent Hints**: AI-generated hints for technical problems
- **Technical Feedback**: Job-context aware feedback with hireability scoring
- **Multi-Language Support**: Python, JavaScript, Java
- **MongoDB Integration**: Persistent session and question storage

## Prerequisites

- Go 1.19+
- MongoDB (local or Atlas)
- Google Gemini API key

## Setup

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Set up environment variables:**
   ```bash
   # Create .env file with your credentials
   MONGODB_URI=mongodb://localhost:27017/stormhacks
   GEMINI_API_KEY=your_gemini_api_key_here
   ```

3. **Run the server:**
   ```bash
   go run main.go
   ```

4. **Access the API:**
   - Web interface: http://localhost:8080
   - Health check: http://localhost:8080/health

## API Endpoints

- `POST /api/interview/session` - Create interview session
- `GET /api/interview-questions` - Get AI-customized questions
- `POST /api/interview/feedback` - Generate interview feedback
- `GET /api/technical-question` - Get technical questions by difficulty
- `POST /api/hint` - Generate AI hints
- `POST /api/execute-code` - Execute and validate code
- `POST /api/technical-feedback` - Generate technical feedback

## Quick Start

1. **Create a session:**
   ```bash
   curl -X POST http://localhost:8080/api/interview/session \
     -H "Content-Type: application/json" \
     -d '{"parsedResumeText": "...", "jobTitle": "Software Engineer", "jobInfo": "..."}'
   ```

2. **Get questions:**
   ```bash
   curl "http://localhost:8080/api/interview-questions?sessionId=YOUR_SESSION_ID"
   ```

3. **Execute code:**
   ```bash
   curl -X POST http://localhost:8080/api/execute-code \
     -H "Content-Type: application/json" \
     -d '{"questionId": "...", "code": "def solution(): ...", "language": "python"}'
   ```

## Project Structure

```
stormhacks-BE/
├── handlers/          # HTTP request handlers
├── services/          # Business logic and AI integration
├── repositories/      # Database operations
├── models/           # Data structures
├── types/            # Request/response types
├── prompts/          # AI prompt templates
├── database/         # MongoDB connection and migrations
└── main.go           # Server entry point
```

## Technologies

- **Backend**: Go with HTTP handlers
- **Database**: MongoDB with BSON
- **AI**: Google Gemini API
- **Code Execution**: go-piston library
- **Architecture**: Clean layered architecture (handlers → services → repositories)