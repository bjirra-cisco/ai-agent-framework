package registry

import (
	"ai-agent-framework/agent"
	"sync"
)

type RegistryCache struct {
	mu     sync.RWMutex
	Agents []agent.Agent
}

func (c *RegistryCache) Set(agents []agent.Agent) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Agents = agents
}

func (c *RegistryCache) Get() []agent.Agent {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.Agents
}
