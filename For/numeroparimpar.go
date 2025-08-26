package main

import "fmt"

func main() {
	for i := 0; i <= 31; i++ {
		if i%2 == 0 {
			fmt.Println(i, "é par")
		} else {
			fmt.Println(i, "é impar")
		}
	}
}
