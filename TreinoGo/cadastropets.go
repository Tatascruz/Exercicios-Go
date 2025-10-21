package main

import "fmt"

type Pet struct {
	Nome  string
	Idade int
	Tipo  string
}

func main() {
	var pets []Pet
	var quantidade int

	fmt.Print("Quantos pets você quer cadastrar? ")
	fmt.Scanln(&quantidade)

	for i := 0; i < quantidade; i++ {
		var p Pet
		fmt.Printf("\nDigite o nome do pet: ")
		fmt.Scanln(&p.Nome)

		fmt.Printf("Digite a idade do pet: ")
		fmt.Scanln(&p.Idade)

		fmt.Printf("Digite o tipo do pet: ")
		fmt.Scanln(&p.Tipo)

		pets = append(pets, p)

	}

	fmt.Println("\n--- Pets cadastrados ---")
	for _, p := range pets {
		fmt.Printf("Nome: %s | Idade: %d | Tipo: %s\n", p.Nome, p.Idade, p.Tipo)
	}

}
