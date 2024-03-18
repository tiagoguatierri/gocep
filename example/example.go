package main

import (
	"fmt"

	"github.com/tiagoguatierri/gocep"
)

func main() {
	cep := gocep.NewCep()
	result := cep.Fetch("01001000")
	fmt.Println(cep.ToJSON(result))
}

// 	{
// 		"zipCode": "01001-000",
// 		"city": "São Paulo",
// 		"state": "SP",
// 		"street": "Praça da Sé",
// 		"district": "Sé",
// 		"provider": "Viacep"
//  }
