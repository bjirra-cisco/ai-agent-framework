package context

import (
	"ai-agent-framework/agent"
	"encoding/json"
)

type SharedContext struct {
	ConversationID string                `json:"conversation_id"`
	MessageID      string                `json:"message_id"`
	UserID         string                `json:"user_id"`
	Prompt         string                `json:"prompt"`
	Params         agent.ExtractedParams `json:"params"`
	Response       json.RawMessage       `json:"response"`
	UIResponse     json.RawMessage       `json:"uiResponse"`
	ResponseSource string                `json:"response_source"`
}
