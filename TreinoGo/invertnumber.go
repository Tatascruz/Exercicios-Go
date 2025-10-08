package main

import "fmt"

func Invert(numbers []int) []int {
	var invertNumbers []int
	for _, number := range numbers {
		invertNumbers = append(invertNumbers, number*(-1))
	}
	return invertNumbers

}

func main() {
	numbers := []int{10, 65, 98, 74, 88, 95, 36}
	result := Invert(numbers)
	fmt.Println("Resultado para o array principal:", result)
}
