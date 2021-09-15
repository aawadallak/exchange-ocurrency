package repository

import (
	"context"
	"exchange_currency/database"
	domain "exchange_currency/domain/ocurrency"
	"exchange_currency/dto"
)

type OcurrencyRepository struct{}

func (d *OcurrencyRepository) Store(dmn *domain.Ocurrency) error {

	db := database.GetDatabase()
	ctx := context.Background()

	tx, err := db.Begin(ctx)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	str := `INSERT INTO exchange_ocurrencies ("from", "to", amount, rate) VALUES ($1, $2, $3, $4)`

	_, err = tx.Exec(ctx, str, dmn.GetFrom(), dmn.GetTo(), dmn.GetAmount(), dmn.GetRate())

	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		tx.Rollback(ctx)
	}

	return nil
}

func (d *OcurrencyRepository) FindByID(id uint) (*domain.Ocurrency, error) {

	db := database.GetDatabase()
	ctx := context.Background()

	str := `SELECT e.id, e.from, e.to, e.amount, e.rate FROM exchange_ocurrencies e where e.id = $1`

	var ocurrency dto.OcurrencyDTO

	err := db.QueryRow(ctx, str, id).Scan(&ocurrency.Id, &ocurrency.From, &ocurrency.To, &ocurrency.Amount, &ocurrency.Rate)

	if err != nil {
		return nil, err
	}

	dmn, err := new(domain.OcurrencyBuilder).
		WithID(ocurrency.Id).
		WithFrom(ocurrency.From).
		WithAmount(ocurrency.Amount).
		WithTo(ocurrency.To).
		WithRate(ocurrency.Rate).
		Ocurrency()

	if err != nil {
		return nil, err
	}

	return dmn, nil
}

func (d *OcurrencyRepository) FindAll() ([]*domain.Ocurrency, error) {

	db := database.GetDatabase()
	ctx := context.Background()

	str := `SELECT e.id, e.from, e.to, e.amount, e.rate FROM exchange_ocurrencies e`

	var ocurrencies []*domain.Ocurrency

	rows, err := db.Query(ctx, str)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	counter := 0

	for rows.Next() {

		counter += 1

		var ocurrency dto.OcurrencyDTO

		err := rows.Scan(&ocurrency.Id, &ocurrency.From, &ocurrency.To, &ocurrency.Amount, &ocurrency.Rate)

		if err != nil {
			return nil, err
		}

		dmn, err := new(domain.OcurrencyBuilder).
			WithID(ocurrency.Id).
			WithFrom(ocurrency.From).
			WithAmount(ocurrency.Amount).
			WithTo(ocurrency.To).
			WithRate(ocurrency.Rate).
			Ocurrency()

		if err != nil {
			return nil, err
		}

		ocurrencies = append(ocurrencies, dmn)
	}

	if counter == 0 {
		return nil, nil
	}

	return ocurrencies, nil
}
