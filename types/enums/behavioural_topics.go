package enums

// BehaviouralTopic represents the available behavioural interview topics
type BehaviouralTopic string

const (
	// BehaviouralTopicGeneral represents general behavioural questions
	BehaviouralTopicGeneral BehaviouralTopic = "General"
	
	// BehaviouralTopicWorkplaceBehavior represents workplace behavior questions
	BehaviouralTopicWorkplaceBehavior BehaviouralTopic = "Workplace Behavior"
	
	// BehaviouralTopicLeadership represents leadership questions
	BehaviouralTopicLeadership BehaviouralTopic = "Leadership"
	
	// BehaviouralTopicProblemSolving represents problem solving questions
	BehaviouralTopicProblemSolving BehaviouralTopic = "Problem Solving"
	
	// BehaviouralTopicConflictResolution represents conflict resolution questions
	BehaviouralTopicConflictResolution BehaviouralTopic = "Conflict Resolution"
	
	// BehaviouralTopicAdaptability represents adaptability questions
	BehaviouralTopicAdaptability BehaviouralTopic = "Adaptability"
	
	// BehaviouralTopicTimeManagement represents time management questions
	BehaviouralTopicTimeManagement BehaviouralTopic = "Time Management"
	
	// BehaviouralTopicCustomerFocus represents customer focus questions
	BehaviouralTopicCustomerFocus BehaviouralTopic = "Customer Focus"
	
	// BehaviouralTopicInnovationCreativity represents innovation & creativity questions
	BehaviouralTopicInnovationCreativity BehaviouralTopic = "Innovation & Creativity"
)

// GetAllBehaviouralTopics returns all available behavioural topics
func GetAllBehaviouralTopics() []BehaviouralTopic {
	return []BehaviouralTopic{
		BehaviouralTopicGeneral,
		BehaviouralTopicWorkplaceBehavior,
		BehaviouralTopicLeadership,
		BehaviouralTopicProblemSolving,
		BehaviouralTopicConflictResolution,
		BehaviouralTopicAdaptability,
		BehaviouralTopicTimeManagement,
		BehaviouralTopicCustomerFocus,
		BehaviouralTopicInnovationCreativity,
	}
}

// IsValidBehaviouralTopic checks if a topic is valid
func IsValidBehaviouralTopic(topic BehaviouralTopic) bool {
	validTopics := GetAllBehaviouralTopics()
	for _, validTopic := range validTopics {
		if validTopic == topic {
			return true
		}
	}
	return false
}
