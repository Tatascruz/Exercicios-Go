package main

import "fmt"

func Well(words []string) string {
	goodCount := 0

	for _, idea := range words {
		if idea == "good" {
			goodCount++

		}
	}

	if goodCount == 0 {
		return "Fail!"
	} else if goodCount <= 2 {
		return "Publish!"
	} else {
		return "I smell a series"
	}

}

func main() {
	test1 := []string{"bad", "good", "bad", "good", "bad"}
	result1 := Well(test1)
	fmt.Println("Teste 1:", result1)

	test2 := []string{"bad", "bad", "bad"}
	result2 := Well(test2)
	fmt.Println("Teste 2:", result2)

	test3 := []string{"good", "good", "good", "bad", "bad"}
	result3 := Well(test3)
	fmt.Println("Teste 3:", result3)

}
