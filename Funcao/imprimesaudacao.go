package main

import "fmt"

func saudacao(nome string) {
	fmt.Println("Olá,", nome, "bem-vindo ao mundo Go!")
}

func main() {
	saudacao("TATA")
	saudacao("ANNA")
	saudacao("THAY")
}
