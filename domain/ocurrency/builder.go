package ocurrency

import "fmt"

type OcurrencyBuilder struct {
	ocurrency Ocurrency
}

func (b *OcurrencyBuilder) WithID(id uint) *OcurrencyBuilder {
	b.ocurrency.id = id
	return b
}

func (b *OcurrencyBuilder) WithAmount(amount float64) *OcurrencyBuilder {
	b.ocurrency.amount = amount
	return b
}

func (b *OcurrencyBuilder) WithFrom(from string) *OcurrencyBuilder {
	b.ocurrency.from = from
	return b
}

func (b *OcurrencyBuilder) WithTo(to string) *OcurrencyBuilder {
	b.ocurrency.to = to
	return b
}

func (b *OcurrencyBuilder) WithRate(rate float64) *OcurrencyBuilder {
	b.ocurrency.rate = rate
	return b
}

func (b *OcurrencyBuilder) Ocurrency() (*Ocurrency, error) {

	if b.ocurrency.from == b.ocurrency.to {
		return nil, fmt.Errorf("não é póssivel realizar conversões entre o mesmo tipo")
	}

	err := validateFromAndTo(b.ocurrency.from, b.ocurrency.to)

	if err != nil {
		return nil, err
	}

	return &b.ocurrency, nil
}
