package main

import "fmt"

func main() {
	numeros := []int{2, 13, 36, 55, 28, 44, 96, 1333, 226, 1356}
	soma := 0

	for _, num := range numeros {
		if num%2 != 0 { // verifica se é par
			soma += num

		}
	}

	fmt.Println("A soma dos números impares é:", soma)

}
