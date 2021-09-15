package ocurrency

func (o *Ocurrency) ToJSON() interface{} {
	return struct {
		ValorConvertido float64 `json:"valorConvertido"`
		SimboloMoeda    string  `json:"simboloMoeda"`
	}{
		ValorConvertido: o.ExchangeOcurrency(),
		SimboloMoeda:    o.CoinSymbol(o.GetTo()),
	}
}

func (o *Ocurrency) DomainToJSON() interface{} {
	return struct {
		MoedaBase       string  `json:"moedaBase"`
		Quantidade      float64 `json:"quantidade"`
		MoedaTroca      string  `json:"moedaTroca"`
		Cambio          float64 `json:"cambio"`
		ValorConvertido float64 `json:"valorConvertido"`
		SimboloMoeda    string  `json:"simboloMoeda"`
	}{
		MoedaBase:       o.from,
		Quantidade:      o.amount,
		MoedaTroca:      o.to,
		Cambio:          o.rate,
		ValorConvertido: o.ExchangeOcurrency(),
		SimboloMoeda:    o.CoinSymbol(o.GetTo()),
	}
}
