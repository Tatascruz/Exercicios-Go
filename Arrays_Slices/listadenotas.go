package main

import "fmt"

func main() {
	notas := []float64{7.5, 9.0, 5.5, 8.0, 4.9, 10.0, 6.8} // Lista de notas

	// Percorrendo a lista de notas
	for i := 0; i < len(notas); i++ {
		if notas[i] >= 7 {
			fmt.Println("Nota", notas[i], "- Aprovado")
		} else if notas[i] >= 5 {
			fmt.Println("Notas", notas[i], "- Recuperação")
		} else {
			fmt.Println("Nota", notas[i], "- Reprovado")
		}

	}

}
