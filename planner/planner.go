package planner

type Plan struct {
	ID             string         `json:"id"`
	ConversationID string         `json:"conversation_id"`
	Steps          []Step         `json:"steps"`
	MetaData       map[string]any `json:"meta,omitempty"`
	Prompt         string         `json:"prompt"`
}

type Step struct {
	StepID     string      `json:"step_id"`
	Tool       string      `json:"tool"`
	Input      interface{} `json:"input"`
	Dependency []string    `json:"dependency"`
	Retry      int         `json:"retry"`
	Priority   int         `json:"priority"`
}

type StepResult struct {
	StepID   string      `json:"step_id"`
	Ok       bool        `json:"ok"`
	Error    string      `json:"error"`
	Data     interface{} `json:"data"`
	Duration int64       `json:"duration"`
}
