package main

import "fmt"

type Pagamento interface {
	Calcular(valor float64) float64
}

type CartaoCredito struct{}
type Dinheiro struct{}

func (c CartaoCredito) Calcular(valor float64) float64 {
	return valor * 1.05
}

func (d Dinheiro) Calcular(valor float64) float64 {
	return valor * 0.90
}

func main() {
	valorCompra := 3256.00

	pagamentos := []Pagamento{
		CartaoCredito{},
		Dinheiro{},
	}

	for _, p := range pagamentos {
		valorFinal := p.Calcular(valorCompra)
		fmt.Printf("Valor final a pagar: R$ %.2f\n", valorFinal)
	}
}
