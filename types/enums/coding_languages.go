package enums

type CodingLanguage string

const (
	CodingLanguagePython     CodingLanguage = "python"
	CodingLanguageJavaScript CodingLanguage = "js"
)

// GetAllCodingLanguages returns all available coding languages
func GetAllCodingLanguages() []CodingLanguage {
	return []CodingLanguage{
		CodingLanguagePython,
		CodingLanguageJavaScript,
	}
}

// IsValidCodingLanguage checks if a language is valid
func IsValidCodingLanguage(language string) bool {
	validLanguages := []string{
		string(CodingLanguagePython),
		string(CodingLanguageJavaScript),
	}
	
	for _, validLang := range validLanguages {
		if language == validLang {
			return true
		}
	}
	return false
}
