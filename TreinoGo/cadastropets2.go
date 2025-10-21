package main

import "fmt"

type Pet struct {
	Nome  string
	Idade int
	Tipo  string
}

func cadastrarPet() Pet {
	var p Pet
	fmt.Print("Digite o nome do Pet: ")
	fmt.Scanln(&p.Nome)

	fmt.Print("Digite a idade do Pet: ")
	fmt.Scanln(&p.Idade)

	fmt.Print("DIgite o tipo do Pet: ")
	fmt.Scanln(&p.Tipo)

	return p
}

func listarPets(pets []Pet) {
	fmt.Println("\n-----------------------")
	fmt.Println("📝 Lista de Pets Cadastrados 📝")
	fmt.Println("-------------------------")

	for _, p := range pets {
		fmt.Printf("Nome: %s | Idade: %d | Tipo: %s\n", p.Nome, p.Idade, p.Tipo)
	}
}

func main() {
	var pets []Pet
	var quantidade int

	fmt.Println("🐾 Bem-vindo(a) ao Cadastro de Pets 🐾")
	fmt.Print("Quantos pets você quer cadastrar? ")
	fmt.Scanln(&quantidade)

	for i := 0; i < quantidade; i++ {
		fmt.Printf("\n🐶 Cadastro do %d pet  😽\n", i+1)
		p := cadastrarPet()
		pets = append(pets, p)
	}

	listarPets(pets)

	fmt.Println("\n🐾 Fim de cadastro! 🐾")
}
