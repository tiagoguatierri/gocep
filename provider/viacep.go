package provider

import (
	"encoding/json"

	"github.com/tiagoguatierri/gocep/model"
)

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

type ViaCep ProviderInstance[ViaCepResponse]

func viaCepParser(data ViaCepResponse) *model.Response {
	return &model.Response{
		ZipCode:  data.Cep,
		City:     data.Localidade,
		State:    data.Uf,
		Street:   data.Logradouro,
		District: data.Bairro,
	}
}

func NewViaCep() *ViaCep {
	p := NewProvider[ViaCepResponse](
		"Viacep",
		"https://viacep.com.br/ws/",
		WithConfig(
			GET,
			nil,
			viaCepParser,
			TIMEOUT,
		),
	)

	viacep := (*ViaCep)(p)
	viacep.super = p
	return viacep
}

func (p *ViaCep) Get(cep string) (*model.Response, error) {
	path := cep + "/json"
	r, err := p.super.Fetch(path, nil)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	var target ViaCepResponse
	err = json.NewDecoder(r.Body).Decode(&target)
	if err != nil {
		return nil, err
	}

	return p.config.parserFn(target), nil
}

func (p *ViaCep) ProviderName() string {
	return p.name
}
