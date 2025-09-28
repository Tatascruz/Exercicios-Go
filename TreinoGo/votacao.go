package main

import "fmt"

func main() {
	var idade int
	fmt.Print("Digite sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		fmt.Println("Você é maior de idade e o voto é obrigatório!")
	} else if idade >= 16 && idade < 18 {
		fmt.Println("Seu voto é opcional.")
	} else {
		fmt.Println("Você é muito jovem e não pode votar!")
	}
}
