package main

import (
	"bufio"
	"fmt"
	"os"
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
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite palavras separadas por espaço:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Split(input, " ")

	allParts := PartList(words)

	fmt.Println("\nDivisões possíveis:\n")
	for _, pair := range allParts {
		fmt.Printf("(%s, %s)\n", pair[0], pair[1])
	}
}
