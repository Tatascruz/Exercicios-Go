package main

import "fmt"

func main() {
	var nota float64 = 5

	if nota >= 7 {
		fmt.Println("Aprovado")
	} else if nota >= 5 {
		fmt.Println("Recuperação")
	} else {
		fmt.Println("Reprovado")
	}

}
