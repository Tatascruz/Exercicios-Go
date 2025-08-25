package main

import "fmt"

func main() {
	var numero int = 4529

	if numero > 0 {
		fmt.Println("Número é positivo")
	} else if numero < 0 {
		fmt.Println("Número é negativo")
	} else {
		fmt.Println("Número é zero")
	}
}
