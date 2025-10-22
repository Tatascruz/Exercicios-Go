package main

import "fmt"

type Pagamento interface {
	Pagar(valor float64)
}

type CartaoCredito struct {
	nome string
}

type Dinheiro struct{}

func (c CartaoCredito) Pagar(valor float64) {
	fmt.Printf("Pagamento de R$ %.2f feito com cartão de crédito (%s)\n", valor, c.nome)
}

func (d Dinheiro) Pagar(valor float64) {
	fmt.Printf("Pagemento de R$ %.2f feito em dinheiro\n", valor)
}

func main() {
	cc := CartaoCredito{nome: "MasterCard"}
	din := Dinheiro{}

	pagamentos := []Pagamento{cc, din}

	for _, p := range pagamentos {
		p.Pagar(350.00)
	}
}
