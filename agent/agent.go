package agent

import (
	"time"
)

type Agent struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Domain       string    `json:"domain"`
	Capabilities []string  `json:"capabilities"`
	Intents      []string  `json:"intents"`
	Priority     int       `json:"priority"`
	Endpoint     string    `json:"endpoint"`
	IsActive     bool      `json:"is_active"`
	LastUpdated  time.Time `json:"last_updated"`
	Config       Config    `json:"config"`
}

type Config struct {
	TimeoutMs  int `json:"timeout_ms"`
	MaxResults int `json:"max_results"`
}
