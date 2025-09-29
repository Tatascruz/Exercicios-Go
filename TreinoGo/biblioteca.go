package main

import "fmt"

type Livro struct {
	Titulo       (string)
	Autor        (string)
	AnoPublicado (int)
}

func main() {
	meuLivro := Livro{Titulo: "A Arte de Programar", Autor: "Tata", AnoPublicado: 2025}
	fmt.Printf("Título: %s\nAutor: %s\nAno de Publicação: %d\n", meuLivro.Titulo, meuLivro.Autor, meuLivro.AnoPublicado)
}
