package registry

import (
	"ai-agent-framework/agent"
	"context"
	"net/http"

	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
)

type RegistryInterface interface {
	Register(agent agent.Agent) error
	GetApplications() ([]agent.Agent, error)
}

type RegistryClient struct {
	config      Config
	context     context.Context
	httpClient  *resty.Client
	oAuthClient *http.Client
}

func NewRegistryClient(ctx context.Context, config Config) *RegistryClient {
	return &RegistryClient{context: ctx, config: config, httpClient: resty.New()}
}

func (r *RegistryClient) Register(agent agent.Agent) error {
	return nil
}

func (r *RegistryClient) GetHttpClient() *resty.Client {
	if r.oAuthClient != nil {

		return r.httpClient.SetTransport(r.oAuthClient.Transport)
	}
	if len(r.config.TokenUrl) < 1 || len(r.config.ClientId) < 1 || len(r.config.ClientSecret) < 1 {
		oAuthConfig := &clientcredentials.Config{
			TokenURL:     r.config.TokenUrl,
			ClientID:     r.config.ClientId,
			ClientSecret: r.config.ClientSecret,
		}
		r.oAuthClient = oAuthConfig.Client(r.context)
		return r.httpClient.SetTransport(r.oAuthClient.Transport)
	}
	return r.httpClient
}
