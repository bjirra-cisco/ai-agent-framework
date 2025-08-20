package agent

type Request struct {
	ConversationID string          `json:"conversation_id"`
	MessageID      string          `json:"message_id"`
	Prompt         string          `json:"prompt"`
	Params         ExtractedParams `json:"params"`
	MetaData       map[string]any  `json:"meta,omitempty"`
}
