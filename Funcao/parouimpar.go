package main

import "fmt"

func parOuImpar(n int) string {
	if n%2 == 0 {
		return "par"
	}
	return "impar"
}

func main() {
	numeros := []int{3, 4, 7, 10, 353, 1426, 947, 29, 3150}

	for i := 0; i < len(numeros); i++ {
		fmt.Println(numeros[i], "é", parOuImpar(numeros[i]))
	}
}
