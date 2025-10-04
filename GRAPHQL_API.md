# GraphQL API Documentation

This backend now uses GraphQL instead of REST endpoints. Here's how to use the API:

## Endpoints

-   **GraphQL Endpoint**: `POST /graphql`
-   **Web Interface**: `GET /` (for testing)

## Available Operations

### Enums

#### BehaviouralTopic

Available behavioural interview topics:

-   `GENERAL` - General behavioural questions
-   `WORKPLACE_BEHAVIOR` - Workplace behavior and professionalism
-   `LEADERSHIP` - Leadership and management skills
-   `PROBLEM_SOLVING` - Problem solving and analytical thinking
-   `CONFLICT_RESOLUTION` - Conflict resolution and communication
-   `ADAPTABILITY` - Adaptability and flexibility
-   `TIME_MANAGEMENT` - Time management and organization
-   `CUSTOMER_FOCUS` - Customer focus and service orientation
-   `INNOVATION_CREATIVITY` - Innovation and creativity

### Queries

#### Get Interview Session

```graphql
query {
    getInterviewSession(sessionId: 1) {
        id
        sessionId
        jobTitle
        companyName
        parsedResumeText
        behaviouralTopics
        createdAt
    }
}
```

#### Generate Interview Questions

```graphql
query {
    generateInterviewQuestions(sessionId: 1) {
        sessionId
        questions {
            question
            hints
        }
    }
}
```

#### Get Behavioural Feedback

```graphql
query {
    getBehaviouralFeedback(sessionId: 1) {
        sessionId
        questionFeedback {
            question
            response
            score
            feedback
            suggestions
        }
    }
}
```

### Mutations

#### Create Interview Session

```graphql
mutation {
    createInterviewSession(
        input: {
            sessionId: 1
            parsedResumeText: "Software Engineer with 5 years experience..."
            jobTitle: "Senior Software Engineer"
            jobInfo: "Full-stack development role..."
            companyName: "Tech Corp"
            additionalInfo: "Remote position"
            behaviouralTopics: ["leadership", "problem-solving"]
        }
    ) {
        id
        sessionId
        jobTitle
        companyName
    }
}
```

#### Submit Behavioural Feedback

```graphql
mutation {
    submitBehaviouralFeedback(
        input: {
            sessionId: 1
            responses: [
                {
                    question: "Tell me about a challenging project"
                    response: "I led a team of 5 developers..."
                }
            ]
        }
    ) {
        sessionId
        questionFeedback {
            question
            response
            score
            feedback
            suggestions
        }
    }
}
```

## Testing with curl

### Create a session:

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation { createInterviewSession(input: { sessionId: 1, parsedResumeText: \"Sample resume\", jobTitle: \"Software Engineer\", jobInfo: \"Full stack development\" }) { sessionId jobTitle } }"
  }'
```

### Get a session:

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{
    "query": "{ getInterviewSession(sessionId: 1) { sessionId jobTitle companyName } }"
  }'
```

## Key Benefits of GraphQL

1. **Single Endpoint**: All operations go through `/graphql`
2. **Flexible Queries**: Request only the data you need
3. **Type Safety**: Built-in validation and documentation
4. **No Over-fetching**: Get exactly what you request
5. **Real-time Schema**: Schema serves as API documentation

## Database

The API uses MongoDB to store interview sessions. Make sure MongoDB is running on `localhost:27017` or update the connection string in the code.
