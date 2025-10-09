package main

import "fmt"

func CalculateYars(humanYears int) (int, int, int) {
	catYears := 0
	dogYears := 0

	if humanYears == 1 {
		catYears = 15
		dogYears = 15
	} else if humanYears == 2 {
		catYears = 24
		dogYears = 24
	} else {
		catYears = 24 + 4*(humanYears-2)
		dogYears = 24 + 5*(humanYears-2)
	}
	return humanYears, catYears, dogYears

}
func main() {
	humanAge := 23
	h, c, d := CalculateYars(humanAge)
	fmt.Println("Idade Humana:", h)
	fmt.Println("Idade de gato:", c)
	fmt.Println("Idade de cachorro:", d)
}
