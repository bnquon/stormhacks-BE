package constants

// BehaviouralTopic represents the available behavioural interview topics
type BehaviouralTopic string

const (
	General              BehaviouralTopic = "general"
	WorkplaceBehavior    BehaviouralTopic = "workplace_behavior"
	Leadership           BehaviouralTopic = "leadership"
	ProblemSolving       BehaviouralTopic = "problem_solving"
	ConflictResolution   BehaviouralTopic = "conflict_resolution"
	Adaptability         BehaviouralTopic = "adaptability"
	TimeManagement       BehaviouralTopic = "time_management"
	CustomerFocus        BehaviouralTopic = "customer_focus"
	InnovationCreativity BehaviouralTopic = "innovation_creativity"
)

// GetAllBehaviouralTopics returns all available behavioural topics
func GetAllBehaviouralTopics() []BehaviouralTopic {
	return []BehaviouralTopic{
		General,
		WorkplaceBehavior,
		Leadership,
		ProblemSolving,
		ConflictResolution,
		Adaptability,
		TimeManagement,
		CustomerFocus,
		InnovationCreativity,
	}
}

// GetBehaviouralTopicStrings returns all available behavioural topics as strings
func GetBehaviouralTopicStrings() []string {
	topics := GetAllBehaviouralTopics()
	strings := make([]string, len(topics))
	for i, topic := range topics {
		strings[i] = string(topic)
	}
	return strings
}

// IsValidBehaviouralTopic checks if a given string is a valid behavioural topic
func IsValidBehaviouralTopic(topic string) bool {
	for _, validTopic := range GetAllBehaviouralTopics() {
		if string(validTopic) == topic {
			return true
		}
	}
	return false
}
