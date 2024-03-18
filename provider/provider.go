package provider

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tiagoguatierri/gocep/model"
)

const (
	GET     Method = "GET"
	POST    Method = "POST"
	TIMEOUT        = 5 * time.Second
)

type Method string

type Provider interface {
	Get(cep string) (*model.Response, error)
	ProviderName() string
}

type ParserFn[T any] func(data T) *model.Response

type ProviderConfig[T any] struct {
	method   Method
	header   http.Header
	parserFn ParserFn[T]
	timeout  time.Duration
}

type ProviderInstance[T any] struct {
	name   string
	url    string
	config *ProviderConfig[T]
	super  *ProviderInstance[T]
}

type ProviderOption[T any] func(p *ProviderInstance[T])

func defaultConfig[T any]() *ProviderConfig[T] {
	return &ProviderConfig[T]{
		method:   GET,
		header:   http.Header{},
		parserFn: nil,
		timeout:  TIMEOUT,
	}
}

func NewProvider[T any](
	name, url string,
	opts ...ProviderOption[T],
) *ProviderInstance[T] {
	p := new(ProviderInstance[T])
	p.name = name
	p.url = url
	p.config = defaultConfig[T]()

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func WithConfig[T any](
	method Method,
	header http.Header,
	parserFn ParserFn[T],
	timeout time.Duration,
) ProviderOption[T] {
	return func(p *ProviderInstance[T]) {
		p.config.method = method
		p.config.header = header
		p.config.parserFn = parserFn
		p.config.timeout = timeout
	}
}

func (p ProviderInstance[T]) Fetch(path string, body io.Reader) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", p.url, path)
	req, err := http.NewRequest(string(p.config.method), url, body)
	if err != nil {
		return nil, err
	}

	req.Header = p.config.header
	client := http.Client{
		Timeout: p.config.timeout,
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
