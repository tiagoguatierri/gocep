package gocep

import (
	"encoding/json"
	"time"

	"github.com/tiagoguatierri/gocep/model"
	"github.com/tiagoguatierri/gocep/provider"
)

type Cep struct {
	providers []provider.Provider
	timeout   time.Duration
}

type CepOption func(c *Cep)

func NewCep(opts ...CepOption) *Cep {
	c := &Cep{
		providers: []provider.Provider{
			provider.NewViaCep(),
			provider.NewBrasilApi(),
		},
		timeout: 5 * time.Second,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func WithProviders(p provider.Provider) CepOption {
	return func(c *Cep) {
		c.providers = append(c.providers, p)
	}
}

func WithTimeout(timeout time.Duration) CepOption {
	return func(c *Cep) {
		c.timeout = timeout
	}
}

func (c *Cep) AddProvider(p provider.Provider) {
	c.providers = append(c.providers, p)
}

func (c *Cep) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

func (c *Cep) ToJSON(result *model.Response) string {
	s, _ := json.MarshalIndent(result, "", "\t")
	return string(s)
}

func (c *Cep) Fetch(cep string) *model.Response {
	result := make(chan *model.Response, len(c.providers))

	for _, p := range c.providers {
		go func(p provider.Provider) {
			resp, err := p.Get(cep)
			if err != nil {
				return
			}

			resp.Provider = p.ProviderName()
			result <- resp
		}(p)
	}

	defer close(result)

	return <-result
}
