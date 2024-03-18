// https://BrasilApi.com.br/api/cep/v1/05545080

package provider

import (
	"encoding/json"

	"github.com/tiagoguatierri/gocep/model"
)

type BrasilApiResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
}

type BrasilApi ProviderInstance[BrasilApiResponse]

func brasilApiParser(data BrasilApiResponse) *model.Response {
	return &model.Response{
		ZipCode:  data.Cep,
		City:     data.City,
		State:    data.State,
		Street:   data.Street,
		District: data.Neighborhood,
	}
}

func NewBrasilApi() *BrasilApi {
	p := NewProvider[BrasilApiResponse](
		"BrasilApi",
		"https://brasilapi.com.br/api/cep/v1",
		WithConfig(
			GET,
			nil,
			brasilApiParser,
			TIMEOUT,
		),
	)

	brasilApi := (*BrasilApi)(p)
	brasilApi.super = p
	return brasilApi
}

func (p *BrasilApi) Get(cep string) (*model.Response, error) {
	r, err := p.super.Fetch(cep, nil)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	var target BrasilApiResponse
	err = json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		return nil, err
	}

	return p.config.parserFn(target), nil
}

func (p *BrasilApi) ProviderName() string {
	return p.name
}
