package config

import (
	"time"

	"github.com/parnurzeal/gorequest"
)

type ClientConfig struct {
	client   *gorequest.SuperAgent
	baseURL  string
	apiKey   string
	username string
}

type IClientConfig interface {
	Client() *gorequest.SuperAgent
	BaseURL() string
	ApiKey() string
	Username() string
}

type Option func(*ClientConfig)

func NewClientConfig(options ...Option) IClientConfig {
	clientConfig := &ClientConfig{
		client: gorequest.New().
			Set("Content-Type", "application/json").
			Set("Accept", "application/json").
			Timeout(30 * time.Second),
	}

	for _, option := range options {
		option(clientConfig)
	}

	return clientConfig
}

func (c *ClientConfig) Client() *gorequest.SuperAgent {
	return c.client
}

func (c *ClientConfig) BaseURL() string {
	return c.baseURL
}

func (c *ClientConfig) ApiKey() string {
	return c.apiKey
}

func (c *ClientConfig) Username() string {
	return c.username
}

// ====== options ======
func WithBaseURL(baseUrl string) Option {
	return func(c *ClientConfig) {
		c.baseURL = baseUrl
	}
}

func WithSignatureKey(signatureKey string) Option {
	return func(c *ClientConfig) {
		c.apiKey = signatureKey
	}
}

func WithUsername(username string) Option {
	return func(c *ClientConfig) {
		c.username = username
	}
}
