package main

import "fmt"

func main() {
	musicas := []string{"Magic shop", "Mikrokosmos", "Serendipity"}

	fmt.Println("Minhas músicas favoritas:")
	for i := 0; i < len(musicas); i++ {
		fmt.Println("-", musicas[i])
	}
}
