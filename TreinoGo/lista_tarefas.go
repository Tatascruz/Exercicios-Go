package main

import "fmt"

func main() {
	tarefas := []string{"Comprar leite", "Pagar contas", "Estudar Go"}
	tarefas = append(tarefas, "Lavar a louça")

	tarefas = append(tarefas[:1], tarefas[1+1:]...)
	fmt.Println(tarefas)
}
