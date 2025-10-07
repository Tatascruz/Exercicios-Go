package main

import "fmt"

func PositiveSum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}

	return sum
}
func main() {
	numbers := []int{1, -4, 7, 12, -25, 48, 32, 11, -19, 23}
	result := PositiveSum(numbers)
	fmt.Println(result)
}
