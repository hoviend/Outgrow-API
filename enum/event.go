package enum

type EventRuleType string

const (
	EventRulesFlat       EventRuleType = "FLAT"
	EventRulesPercentage EventRuleType = "PERCENTAGE"
)
