package main

import "fmt"

func main() {
	var nota1 float64 = 5.0
	var nota2 float64 = 2.5
	var nota3 float64 = 7.5

	var soma float64 = nota1 + nota2 + nota3
	media := soma / 3.0

	fmt.Printf("A média das notas é: %.2f", media)

}
