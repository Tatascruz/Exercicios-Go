package main

import "fmt"

func CountPositiveSumNegatives(numbers []int) []int {
	var qntPositives int
	var sumNegatives int
	for _, number := range numbers {
		if number > 0 {
			qntPositives += 1
		}
		if number < 0 {
			sumNegatives += number
		}
	}

	return []int{qntPositives, sumNegatives}

}
func main() {
	numbers := []int{1, 2, 67, 4, 32, 6, 7, 13, 309, 10, -11, -12, -13, -1554, -15, 50, 41, 99, -87}
	result := CountPositiveSumNegatives(numbers)
	fmt.Println("Resultado para o array principal:", result)
}
