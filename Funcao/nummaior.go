package main

import "fmt"

func maior(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	resultado := maior(510, 1329)
	fmt.Println("O maior número é:", resultado)
}
