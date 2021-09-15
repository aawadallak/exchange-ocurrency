package ocurrency

import "fmt"

func (o *Ocurrency) GetAmount() float64 {
	return o.amount
}

func (o *Ocurrency) GetFrom() string {
	return o.from
}

func (o *Ocurrency) GetTo() string {
	return o.to
}

func (o *Ocurrency) GetRate() float64 {
	return o.rate
}

func (o *Ocurrency) ExchangeOcurrency() float64 {
	return (o.amount * o.rate)
}

func (o *Ocurrency) CoinSymbol(ocurrency string) string {
	switch ocurrency {
	case "BRL":
		return "R$"
	case "USD":
		return "$"
	case "EUR":
		return "€"
	case "BTC":
		return "₿"
	default:
		return ""
	}
}

func validateFromAndTo(from, to string) error {
	if from != "BRL" && from != "BTC" && from != "USD" && from != "EUR" {
		return fmt.Errorf("moéda não disponível em nosso sistema")
	}

	if to != "BRL" && to != "BTC" && to != "USD" && to != "EUR" {
		return fmt.Errorf("moéda não disponível em nosso sistema")
	}
	return nil
}
