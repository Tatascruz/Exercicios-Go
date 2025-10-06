package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"I", "wish", "I", "hadn't", "come"}
	var finalWord string

	for i := 1; i < len(words); i++ {

		word := fmt.Sprintf("(%s, %s)", strings.Join(words[:i], " "), strings.Join(words[i:], " "))
		finalWord += word
	}

	fmt.Println(finalWord)
}
