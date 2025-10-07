package main

import (
	"fmt"
	"strconv"
)

func StringToNumber(str string) int {
	number, _ := strconv.Atoi(str)

	return number
}
func main() {
	num1 := StringToNumber("123456")
	num2 := StringToNumber("15228624")
	num3 := StringToNumber("65897426")

	fmt.Printf("Número 1: %d\n", num1)
	fmt.Printf("Número 2: %d\n", num2)
	fmt.Printf("Número 3: %d\n", num3)

}
