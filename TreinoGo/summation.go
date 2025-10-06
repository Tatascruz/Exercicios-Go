package main

import "fmt"

func Summation(num int) int {
	var sum int = 0

	for i := 1; i <= num; i++ {
		sum += i
	}

	return sum
}

func main() {
	result1 := Summation(2)
	fmt.Println("A soma de 1 até 2 é:", result1)

	result2 := Summation(38)
	fmt.Println("A soma de 1 até 38 é:", result2)

	result3 := Summation(113)
	fmt.Println("A soma de 1 até 113 é:", result3)

}
