package main

import "fmt"

func SortMyString(word string) string {
	var evenLetters, oddLetters string

	for i := 0; i < len(word); i++ {
		if i%2 == 0 {
			evenLetters += string(word[i])
			continue
		}
		oddLetters += string(word[i])
	}
	return fmt.Sprintf("%s %s", evenLetters, oddLetters)

}
func main() {
	word := ("TataSiqueiraCruz")
	result := SortMyString(word)
	fmt.Println(result)

}
