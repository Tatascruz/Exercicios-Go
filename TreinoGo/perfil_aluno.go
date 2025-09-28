package main

import "fmt"

func main() {
	perfil := map[string]string{
		"nome":   "João da Silva",
		"idade":  "25",
		"cidade": "São Paulo",
	}

	perfil["curso"] = "Engenharia de Software"
	perfil["idade"] = "26"

	for chave, valor := range perfil {
		fmt.Printf("Propriedade: %s, Valor: %s\n", chave, valor)

	}
}
