package prompts

import "fmt"

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
3. First Priority: Make the questions more specific to the job role and company
4. Second Priority: Reference relevant technologies, skills, or experiences from the resume when appropriate (do it moderately, so like 1 question out of 3)
5. Maintain the behavioral interview format (STAR method applicable)
6. Keep questions professional and fair
7. KEEP QUESTIONS CONCISE AND TO THE POINT, MAKE THEM SHORT AND TO THE POINT

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

// HintGenerationPrompt creates a prompt for generating interview hints
func HintGenerationPrompt(question string, userCode string, userSpeech string, previousHints []string) string {
	// Build previous hints text
	previousHintsText := ""
	if len(previousHints) > 0 {
		previousHintsText = "\n\nPREVIOUS HINTS GIVEN:\n"
		for i, hint := range previousHints {
			previousHintsText += fmt.Sprintf("%d. %s\n", i+1, hint)
		}
	}

	return `You are an expert technical interviewer conducting a coding interview. Your task is to provide helpful hints to guide the candidate toward solving the problem.

CURRENT INTERVIEW SITUATION:
- Question: ` + question + `
- What the candidate said: ` + userSpeech + `
- Candidate's Current Code: ` + userCode + `"` + previousHintsText + `

NOTE: Pay attention to both what the candidate said. If either contains phrases like "give me the solution", "show me the answer", "I need the full solution", "just tell me how to do it", or similar requests for the complete answer, they are asking for a solution.

CRITICAL INSTRUCTIONS:
- Act as an interviewer trying to guide the interviewee to the solution
- NEVER provide the complete solution or full answer
- Give progressive hints that lead them in the right direction
- Focus on helping them think through the problem step by step
- Be supportive and encouraging but don't give away the answer
- IMPORTANT: If the user asks for a solution but has received fewer than 3 hints, calm them down and continue guiding them instead of providing the solution
- Only provide about 90% of the solution (leaving some details for them to figure out) if they've already received 3 or more hints

IMPORTANT RULES:
- Do NOT repeat any of the previous hints already given
- Count the number of previous hints provided
- If user asks for solution but has fewer than 3 hints: Calm them down with encouraging words like "Don't worry, you're doing great!" and continue guiding them
- If user asks for solution and has 3+ hints: Provide about 90% of the solution approach
- Provide one conversational hint that an interviewer would say out loud
- Provide one concise summary hint for display
- Make hints specific and actionable
- Guide them toward the next logical step in their solution

Return your response in this exact JSON format:
{
  "conversationalHint": "A natural, conversational hint that an interviewer would say out loud to guide the candidate",
  "hintSummary": "A concise summary hint for display purposes"
}

Return ONLY the JSON, no other text.`
}

// TechnicalFeedbackPrompt creates a prompt for generating technical feedback
func TechnicalFeedbackPrompt(questionInfo map[string]string, userCode string, hintsUsed int, isCompleted bool, timeTaken int) string {
	return `You are an expert technical interviewer evaluating a candidate's performance on a coding problem.

JOB CONTEXT:
- Job Title: ` + questionInfo["jobTitle"] + `
- Company: ` + questionInfo["companyName"] + `

PROBLEM INFORMATION:
- Question: ` + questionInfo["question"] + `
- Description: ` + questionInfo["description"] + `
- Difficulty: ` + questionInfo["difficulty"] + `

CANDIDATE PERFORMANCE:
- Code Submitted: ` + userCode + `
- Hints Used: ` + fmt.Sprintf("%d", hintsUsed) + `
- Completed: ` + fmt.Sprintf("%t", isCompleted) + `
- Time Taken: ` + fmt.Sprintf("%d seconds", timeTaken) + `

EVALUATION CRITERIA (adjusted for job level):
- Code correctness and efficiency
- Problem-solving approach
- Code quality and readability
- Time management
- Independence (hints used)
- Seniority expectations based on job title

IMPORTANT: Adjust your evaluation based on the job title:
- For Senior/Lead roles: Expect advanced algorithms, clean architecture, optimal solutions
- For Mid-level roles: Expect solid fundamentals, good problem-solving, some optimization
- For Junior/Entry roles: Focus on basic correctness, learning potential, growth mindset
- For Intern/Co-op roles: Emphasize learning, basic understanding, willingness to improve

Provide feedback in this exact JSON format:
{
  "hireAbilityScore": number (0-100),
  "suggestions": [
    "suggestion 1",
    "suggestion 2",
    "suggestion 3"
  ],
  "strengths": [
    "strength 1",
    "strength 2",
    "strength 3"
  ]
}

Return ONLY the JSON, no other text with the removed beginning and ending quote and json markers`
}
