package dto

import domain "exchange_currency/domain/ocurrency"

type OcurrencyDTO struct {
	Id     uint    `json:"id,omitempty"`
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Rate   float64 `json:"rate"`
}

func (o *OcurrencyDTO) Convert2Entity() (*domain.Ocurrency, error) {

	dmn, err := new(domain.OcurrencyBuilder).
		WithAmount(o.Amount).
		WithFrom(o.From).
		WithRate(o.Rate).
		WithTo(o.To).
		Ocurrency()

	if err != nil {
		return nil, err
	}

	return dmn, nil
}
