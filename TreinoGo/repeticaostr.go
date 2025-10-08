package main

import "fmt"

func repeatStr(n int, word string) string {
	result := ""
	for i := 0; i < n; i++ {
		result += word
	}

	return result

}

func main() {
	word := repeatStr(50, "BTS")

	fmt.Printf(word)

}
