package model

type Response struct {
	ZipCode  string `json:"zipCode"`
	City     string `json:"city"`
	State    string `json:"state"`
	Street   string `json:"street"`
	District string `json:"district"`
	Provider string `json:"provider"`
}
