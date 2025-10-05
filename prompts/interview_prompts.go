package prompts

// QuestionCustomizationPrompt creates a prompt for customizing interview questions
func QuestionCustomizationPrompt(sessionInfo map[string]string, questionsText string) string {
	return `You are an expert interview coach. I need you to customize these behavioral interview questions to be more specific to the candidate's background and the job they're applying for.

JOB INFORMATION:
- Job Title: ` + sessionInfo["jobTitle"] + `
- Job Description: ` + sessionInfo["jobInfo"] + `
- Company: ` + sessionInfo["companyName"] + `
- Additional Info: ` + sessionInfo["additionalInfo"] + `

CANDIDATE BACKGROUND:
- Resume Text: ` + sessionInfo["resumeText"] + `

ORIGINAL QUESTIONS TO CUSTOMIZE:
` + questionsText + `

INSTRUCTIONS:
1. IMPORTANT: Return exactly the same number of questions as provided in the input (do not add or remove questions)
2. Keep the same behavioral topic for each question
3. Make the questions more specific to the job role and company
4. Reference relevant technologies, skills, or experiences from the resume when appropriate
5. Maintain the behavioral interview format (STAR method applicable)
6. Keep questions professional and challenging but fair, but not too long, make it concise and to the point

Please return the customized questions in this exact JSON format:
{
  "questions": [
    {
      "behavioralTopic": "Leadership",
      "question": "Customized question text here",
      "hints": [
        "What specific actions did you take?",
        "Who was involved in the situation?",
        "What was the outcome or result?"
      ]
    }
  ]
}

The hints should be follow-up questions that help the candidate provide a complete STAR response. Each question should have exactly 3 hints that guide them to explain:
- Situation/Task: Context and what needed to be done
- Action: Specific steps they took  
- Result: Outcome and what they learned

Example hints format and length:
- "What was the specific situation or challenge you faced?"
- "What actions did you take to address this challenge?"
- "What was the outcome and what did you learn from this experience?"

KEEP the hint super short and concise, just few words
Return ONLY the JSON, no other text.`
}

// FeedbackEvaluationPrompt creates a prompt for evaluating interview responses
func FeedbackEvaluationPrompt(sessionInfo map[string]string, questionsWithAnswers string) string {
	return `You are an expert interview coach and hiring manager. Evaluate these interview responses based on the candidate's background and the job requirements.

JOB INFORMATION:
- Job Title: ` + sessionInfo["jobTitle"] + `
- Job Description: ` + sessionInfo["jobInfo"] + `
- Company: ` + sessionInfo["companyName"] + `
- Additional Info: ` + sessionInfo["additionalInfo"] + `

CANDIDATE BACKGROUND:
- Resume Text: ` + sessionInfo["resumeText"] + `

INTERVIEW RESPONSES:
` + questionsWithAnswers + `

EVALUATION CRITERIA:
1. Use of STAR method (Situation, Task, Action, Result)
2. Relevance to the role and company
3. Demonstration of required skills and competencies
4. Specificity and detail in responses
5. Leadership and problem-solving capabilities
6. Communication clarity and structure

Please evaluate each response and provide:
- Score (1-10 scale)
- 3 specific strengths
- 3 areas for improvement
- Overall hireability score (0-100)
- 3 points of overall feedback

Return your evaluation in this exact JSON format:
{
  "interviewQuestionFeedback": [
    {
      "question": "The exact interview question",
      "score": number (1-10),
      "strengths": ["strength 1", "strength 2", "strength 3"],
      "areasForImprovement": ["improvement 1", "improvement 2", "improvement 3"],
    }
  ],
  "hireAbilityScore": number (0-100),
  "overallFeedback": ["feedback 1", "feedback 2", "feedback 3"]
}

Return ONLY the JSON, no other text.`
}
