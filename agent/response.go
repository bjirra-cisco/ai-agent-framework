package agent

import "encoding/json"

type Response struct {
	ConversationID string          `json:"conversation_id"`
	MessageID      string          `json:"message_id"`
	Source         string          `json:"source"`
	Endpoint       string          `json:"endpoint"`
	OK             bool            `json:"ok"`
	Status         int             `json:"status"`
	DurationMs     int64           `json:"duration_ms"`
	Data           json.RawMessage `json:"data"`
	Error          *ErrDetail      `json:"error,omitempty"`
	MetaData       map[string]any  `json:"meta,omitempty"`
}

type ErrDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
