package registry

import "time"

type Config struct {
	AddressUrl   string
	Protocol     string //"http" or "grpc"
	PollInterval time.Duration
	TokenUrl     string
	ClientId     string
	ClientSecret string
}
