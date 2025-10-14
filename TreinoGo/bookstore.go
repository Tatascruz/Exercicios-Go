package main

import "fmt"

type Livro struct {
	Titulo       string
	Autor        string
	AnoPublicado int
	Preco        float64
}

func (l Livro) ExibirDetalhes() {
	fmt.Println("---------------------------")
	fmt.Printf("Detalhe do Livro: %s\n", l.Titulo)
	fmt.Printf("Autor: %s\n", l.Autor)
	fmt.Printf("Ano de Publicação: %d\n", l.AnoPublicado)
	fmt.Printf("Preço: R$ %.2f\n", l.Preco)
}

func main() {
	meuLivro := Livro{
		Titulo:       "Harry Potter e a Pedra Filosofal",
		Autor:        "J.K. Rowling",
		AnoPublicado: 1997,
		Preco:        32.30,
	}

	outroLivro := Livro{
		Titulo:       "Crepúsculo",
		Autor:        "Stephenie Meyer",
		AnoPublicado: 2005,
		Preco:        42.24,
	}

	fmt.Println("\n-- Exibindo os livros --")
	meuLivro.ExibirDetalhes()
	outroLivro.ExibirDetalhes()

	fmt.Println("\n== Comparação de Antiguidade ==")

	if meuLivro.AnoPublicado < outroLivro.AnoPublicado {
		fmt.Printf("O livro main antigo é '%s', publicado em %d.\n", meuLivro.Titulo, meuLivro.AnoPublicado)
	} else if outroLivro.AnoPublicado < meuLivro.AnoPublicado {
		fmt.Printf("O livro mais antigo é '%s', publicado em %d.\n", outroLivro.Titulo, outroLivro.AnoPublicado)
	} else {
		fmt.Println("Os livros foram publiacdos no mesmo ano.")

	}

}
