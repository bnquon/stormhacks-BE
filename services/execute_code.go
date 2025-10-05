package services

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	piston "github.com/milindmadhukar/go-piston"
	"stormhacks-be/repositories"
	"stormhacks-be/types/enums"
	"stormhacks-be/types/requests"
	"stormhacks-be/types/responses"
)

func ExecuteCode(input requests.ExecuteTechnicalInput, interviewRepo *repositories.InterviewRepository) (*responses.ExecuteTechnicalResponse, error) {
	// Validate language
	if !enums.IsValidCodingLanguage(string(input.Language)) {
		return nil, fmt.Errorf("invalid language: %s. Allowed languages: python, javascript", input.Language)
	}

	// Get the technical question by ID
	question, err := interviewRepo.GetTechnicalQuestionByID(input.QuestionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get question: %w", err)
	}

	// Execute code for each test case
	var allOutputs []string
	var allErrors []string
	var totalExecutionTime int64
	success := true

	for i, testCase := range question.Question.TestCases {
		// Prepare the code with test case input
		executionCode := prepareCodeWithTestCase(input.Code, testCase.Input, string(input.Language), question.Question.FunctionName)

		fmt.Println(executionCode);
		
		// Execute the code
		startTime := time.Now()
		output, err := executeCodeWithPiston(executionCode, string(input.Language))
		executionTime := time.Since(startTime).Milliseconds()
		totalExecutionTime += executionTime

		if err != nil {
			allErrors = append(allErrors, fmt.Sprintf("Test case %d: %v", i+1, err))
			success = false
			continue
		}

		// Clean and compare output
		cleanedOutput := cleanOutput(output)
		cleanedExpected := cleanOutput(testCase.ExpectedOutput)
		
		allOutputs = append(allOutputs, cleanedOutput)
		
		// Check if output matches expected
		if !compareOutputs(cleanedOutput, cleanedExpected) {
			allErrors = append(allErrors, fmt.Sprintf("Test case %d: Expected '%s', got '%s'", i+1, cleanedExpected, cleanedOutput))
			success = false
		}
	}

	// Determine final output and error
	var finalOutput string
	var finalError string

	if success {
		finalOutput = strings.Join(allOutputs, "\n")
	} else {
		// Categorize and format errors
		finalError = categorizeAndFormatErrors(allErrors, allOutputs)
		finalOutput = strings.Join(allOutputs, "\n")
	}

	return &responses.ExecuteTechnicalResponse{
		QuestionID:    input.QuestionID,
		Code:         input.Code,
		Language:     string(input.Language),
		Output:       finalOutput,
		Error:        finalError,
		ExecutionTime: totalExecutionTime,
		Success:      success,
	}, nil
}

// prepareCodeWithTestCase prepares code with test case input
func prepareCodeWithTestCase(code, testInput, language, functionName string) string {
	switch language {
	case string(enums.CodingLanguagePython):
		return fmt.Sprintf(`
%s

# Test with input: %s
print(%s(%s))
`, code, testInput, functionName, testInput)
	case string(enums.CodingLanguageJavaScript):
		return fmt.Sprintf(`
%s

// Test with input: %s
console.log(%s(%s));
`, code, testInput, functionName, testInput)
	default:
		return fmt.Sprintf(`
%s

# Test with input: %s
print(%s(%s))
`, code, testInput, functionName, testInput)
	}
}

// executeCodeWithPiston executes code using the piston client
func executeCodeWithPiston(code, language string) (string, error) {
	client := piston.CreateDefaultClient()
	
	// Map language to piston language code
	languageCode := mapLanguageToPistonCode(language)
	
	result, err := client.Execute(languageCode, "", []piston.Code{
		{Content: code},
	})
	if err != nil {
		return "", err
	}
	
	return result.GetOutput(), nil
}

// mapLanguageToPistonCode maps language names to piston language codes
func mapLanguageToPistonCode(language string) string {
	languageMap := map[string]string{
		string(enums.CodingLanguagePython):     "python3",
		string(enums.CodingLanguageJavaScript): "js",
	}
	
	if code, exists := languageMap[strings.ToLower(language)]; exists {
		return code
	}
	return "python3" // default to Python
}

// cleanOutput cleans the output for comparison
func cleanOutput(output string) string {
	// Remove extra whitespace and newlines
	cleaned := strings.TrimSpace(output)
	// Remove common prefixes like "Result:" or "Output:"
	cleaned = regexp.MustCompile(`^(Result:|Output:)\s*`).ReplaceAllString(cleaned, "")
	cleaned = strings.TrimSpace(cleaned)
	
	// Normalize array formatting - remove spaces around commas and brackets
	cleaned = regexp.MustCompile(`\[\s*`).ReplaceAllString(cleaned, "[")
	cleaned = regexp.MustCompile(`\s*\]`).ReplaceAllString(cleaned, "]")
	cleaned = regexp.MustCompile(`\s*,\s*`).ReplaceAllString(cleaned, ",")
	
	return cleaned
}

// compareOutputs compares two outputs for equality
func compareOutputs(actual, expected string) bool {
	// Normalize both strings for comparison
	actual = strings.TrimSpace(strings.ToLower(actual))
	expected = strings.TrimSpace(strings.ToLower(expected))
	
	// Direct comparison
	if actual == expected {
		return true
	}
	
	// Try to extract just the result value (remove any extra text)
	actualValue := extractResultValue(actual)
	expectedValue := extractResultValue(expected)
	
	return actualValue == expectedValue
}

// extractResultValue extracts the actual result value from output
func extractResultValue(output string) string {
	// Try to extract the last meaningful value from the output
	lines := strings.Split(output, "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line != "" && !strings.Contains(line, "Result:") && !strings.Contains(line, "Output:") {
			// Extract just the value part
			value := regexp.MustCompile(`.*?(\d+|true|false|"[^"]*"|'[^']*'|\[.*\]|\{.*\})`).FindString(line)
			if value != "" {
				return strings.TrimSpace(value)
			}
		}
	}
	return strings.TrimSpace(output)
}

// categorizeAndFormatErrors categorizes errors and formats them appropriately
func categorizeAndFormatErrors(errors []string, outputs []string) string {
	// Check for compilation errors (SyntaxError, etc.)
	for _, err := range errors {
		if strings.Contains(err, "SyntaxError") || strings.Contains(err, "IndentationError") || 
		  	strings.Contains(err, "NameError") || strings.Contains(err, "TypeError") {
			return "Compilation Error: " + extractMainError(err)
		}
	}
	
	// Check for runtime errors
	for _, err := range errors {
		if strings.Contains(err, "IndexError") || strings.Contains(err, "KeyError") || 
			strings.Contains(err, "AttributeError") || strings.Contains(err, "ValueError") ||
		  strings.Contains(err, "RuntimeError") {
			return "Runtime Error: " + extractMainError(err)
		}
	}
	
	// Check for timeout or memory errors
	for _, err := range errors {
		if strings.Contains(err, "timeout") || strings.Contains(err, "memory") || 
		  strings.Contains(err, "killed") {
			return "Execution Error: " + extractMainError(err)
		}
	}
	
	// If it's just wrong answers, format them nicely
	if len(errors) > 0 && strings.Contains(errors[0], "Expected") {
		return strings.Join(errors, "\n")
	}
	
	// General error fallback
	return "Execution Error: " + strings.Join(errors, "\n")
}

// extractMainError extracts the main error message from a complex error string
func extractMainError(errorStr string) string {
	// Look for common error patterns
	if strings.Contains(errorStr, "SyntaxError:") {
		parts := strings.Split(errorStr, "SyntaxError:")
		if len(parts) > 1 {
			return "SyntaxError:" + parts[1]
		}
	}
	
	if strings.Contains(errorStr, "NameError:") {
		parts := strings.Split(errorStr, "NameError:")
		if len(parts) > 1 {
			return "NameError:" + parts[1]
		}
	}
	
	if strings.Contains(errorStr, "IndexError:") {
		parts := strings.Split(errorStr, "IndexError:")
		if len(parts) > 1 {
			return "IndexError:" + parts[1]
		}
	}
	
	// Return the first line of the error if it's a multi-line error
	lines := strings.Split(errorStr, "\n")
	if len(lines) > 0 {
		return strings.TrimSpace(lines[0])
	}
	
	return errorStr
}
