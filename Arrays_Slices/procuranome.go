package main

import "fmt"

func main() {
	nomes := []string{"Anna", "Tete", "Jk", "Jimin"}
	procurar := "Jimin"
	encontrado := false

	for i := 0; i < len(nomes); i++ {
		if nomes[i] == procurar {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println(procurar, "foi encontrado")
	} else {
		fmt.Println(procurar, "não está na lista")
	}

}
