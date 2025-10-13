package main

import (
	"fmt"
	"strings"
)

func Disemvowel(comment string) string {
	vogais := "aeiouAEIOU"
	var result strings.Builder

	for _, ch := range comment {
		if !strings.ContainsRune(vogais, ch) {
			result.WriteRune(ch)
		}
	}

	return result.String()
}

func main() {

	fmt.Println(Disemvowel("Golang is better!!!"))
}
