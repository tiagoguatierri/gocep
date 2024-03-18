<p  align="center">

<img  src="http://piskel-imgstore-b.appspot.com/img/d580e96e-bd8a-11e6-b157-9949cad4d609.gif">

</p>

<h1  align="center">Go-cep</h1>
</p>

Busca concorrente de CEP diretamente nos serviços da [Viacep](https://viacep.com.br/) e [BrasilAPI](https://brasilapi.com.br/)
 
 Essa lib está sendo construída totalmente baseada na lib [cep-promise](cep-promise)

## Features
* Sempre atualizado em tempo-real por se conectar diretamente aos serviços da ViaCEP e BrasilAPI.

* Possui alta disponibilidade por usar vários serviços como fallback.

* Sempre retorna a resposta mais rápida por fazer as consultas de forma concorrente.

* Sem limites de uso (rate limits) conhecidos.

## Como utilizar
### Realizando uma consulta

``` go
package  main

import (
"fmt"

"github.com/tiagoguatierri/gocep"
)

func  main() {
	cep := gocep.NewCep()
	result := cep.Fetch("01001000")
	fmt.Println(cep.ToJSON(result))
}

// {
// 		"zipCode": "01001-000",
// 		"city": "São Paulo",
// 		"state": "SP",
// 		"street": "Praça da Sé",
// 		"district": "Sé",
// 		"provider": "Viacep"
// }
```  

## Próximos passos
- Validação do cep
- Tratativa de erros por provider
- Adicionar mais providers [https://brasilaberto.com/docs/v1/zipcode, https://opencep.com/v1/15050305, https://cdn.apicep.com/file/apicep]
- Testes unitários
 
## Como contribuir
-	Clone esse repositório
-	Abra uma PR a partir de main
-	Atualize o `changelog`