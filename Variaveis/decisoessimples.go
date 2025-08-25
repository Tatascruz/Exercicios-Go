package main

import "fmt"

func main() {
	var idade int = 30

	if idade >= 18 {
		fmt.Println("Você é maior de idade!")
	} else if idade >= 13 && idade <= 17 {
		fmt.Println("Você é adolescente!")
	} else {
		fmt.Println("Você é menor de idade!")
	}
}
