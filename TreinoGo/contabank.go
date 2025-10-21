package main

import "fmt"

type Conta struct {
	Titular string
	Numero  int
	Saldo   float64
}

func (c *Conta) Depositar(valor float64) {
	if valor > 0 {
		c.Saldo += valor
		fmt.Printf("Depósito de R$ %.2f realizado com sucesso.\n", valor)
	} else {
		fmt.Println("Erro: O valor do depósito deve ser positivo")
	}
}

func (c *Conta) Sacar(valor float64) bool {
	if valor <= c.Saldo && valor > 0 {
		c.Saldo -= valor
		fmt.Printf("Saque de R$ %.2f realizado com sucesso.\n", valor)
		return true
	}
	fmt.Println("Erro: Saque não realizado. Saldo insufuciente ou valor inválido")
	return false
}

func main() {
	minhaConta := Conta{
		Titular: "Alice da Silva",
		Numero:  123456,
		Saldo:   152.63,
	}

	fmt.Println("--- Saldo Inicial ---")
	fmt.Printf("Conta %d (Titular: %s): R$ %.2f\n", minhaConta.Numero, minhaConta.Titular, minhaConta.Saldo)

	fmt.Println("\n--- Realizando Operações ---")
	minhaConta.Depositar(100.00)
	minhaConta.Sacar(50.00)

	fmt.Println("\n--- Saldo Final ---")
	fmt.Printf("Conta %d (Titular: %s) : R$ %.2f\n", minhaConta.Numero, minhaConta.Titular, minhaConta.Saldo)

}
