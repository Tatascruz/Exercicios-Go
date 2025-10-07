package main

import (
	"fmt"
	"strconv"
)

func NumberToString(number int) string {
	s := strconv.Itoa(number)
	return s
}

func main() {
	num1 := NumberToString(123456)
	num2 := NumberToString(15228624)
	num3 := NumberToString(65897426)

	fmt.Printf("Valor como string: %q\n", num1)
	fmt.Printf("Valor como string: %q\n", num2)
	fmt.Printf("Valor como string: %q\n", num3)
}
