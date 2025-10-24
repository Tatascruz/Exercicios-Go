package main

import "fmt"

// Interface com retorno
type Pagamento interface {
	Calcular(valor float64) float64
}

// Tipos concretos
type CartaoCredito struct{}
type Dinheiro struct{}

// Implementações dos métodos
func (c CartaoCredito) Calcular(valor float64) float64 {
	return valor * 1.05 // 5% de desconto
}

func (d Dinheiro) Calcular(valor float64) float64 {
	return valor * 0.90 // 10% de desconto
}

// Função principal
func main() {
	var valor float64
	var opcao int

	fmt.Print("Digite o valor da compra: R$ ")
	fmt.Scan(&valor)

	fmt.Println("Escolha o tipo de pagamento:")
	fmt.Println("1 - Cartão de Crédito (5% de taxa)")
	fmt.Println("2 - Dinheiro (10% de desconto)")
	fmt.Print("Opção: ")
	fmt.Scan(&opcao)

	var p Pagamento

	// Escolha do tipo de pagamento
	switch opcao {
	case 1:
		p = CartaoCredito{}
	case 2:
		p = Dinheiro{}
	default:
		fmt.Println("Opção inválida!")
		return
	}

	valorFinal := p.Calcular(valor)
	fmt.Printf("Valor final a pagar: R$ %.2f\n", valorFinal)
}
