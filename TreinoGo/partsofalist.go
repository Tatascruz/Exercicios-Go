package main

import (
	"fmt"
	"strings"
)

func PartList(words []string) [][]string {
	var result [][]string

	for i := 1; i < len(words); i++ {
		part1 := strings.Join(words[:i], " ")
		part2 := strings.Join(words[i:], " ")
		pair := []string{part1, part2}
		result = append(result, pair)

	}

	return result
}

func main() {
	testList := []string{"I", "Number", "Kiss", "Dog", "Bless", "Age"}
	allParts := PartList(testList)

	fmt.Println("Divisões possíveis:\n")
	for _, pair := range allParts {
		fmt.Printf("[%s, %s]\n", pair[0], pair[1])
	}

}
