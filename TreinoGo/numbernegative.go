package main

import "fmt"

func makeNegative(num int) int {
	if num < 0 {
		return num
	} else {
		num = num * (-1)
		return num
	}
}
func main() {
	fmt.Println(makeNegative(5))
	fmt.Println(makeNegative(-3))
	fmt.Println(makeNegative(0))
}
