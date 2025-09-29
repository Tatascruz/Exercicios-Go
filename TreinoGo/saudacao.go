package main

import "fmt"

func saudar(nome string) string {
	var mensagem = "Olá," + nome + ",bem-vinda ao Go!"
	return mensagem
}

func main() {
	var nomeUsuario = "Tata"
	saudacaoFinal := saudar(nomeUsuario)
	fmt.Println(saudacaoFinal)
}
