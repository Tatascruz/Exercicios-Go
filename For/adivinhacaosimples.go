package main

import "fmt"

func main() {
	var resposta int = 7 // número secreto
	var tentativa int

	fmt.Println("=== Jogo da Adivinhação ===")

	for i := 1; i <= 3; i++ {
		fmt.Println("Tentativa ", i, ": Adivinhe o número (0 a 10): ")
		fmt.Scan(&tentativa)

		if tentativa == resposta {
			fmt.Println("Parabéns! Você acertou!")
			break
		} else {
			fmt.Println("Errado. Tente novamente.")
		}
		if i == 3 {
			fmt.Println("Fim de jogo! O número era:", resposta)
		}

	}
}
