package repository

import (
	"exchange_currency/database"
	"exchange_currency/dto"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const (
	correct_id = 1
	wrong_id   = 0
)

func TestOcurrencyRepository_Store(t *testing.T) {

	if err := godotenv.Load("../../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	database.Init()

	svc := OcurrencyRepository{}

	t.Run("Shouldn't return an error", func(t *testing.T) {

		d := dto.OcurrencyDTO{
			Amount: 10,
			From:   "BRL",
			To:     "USD",
			Rate:   4.3,
		}

		dmn, err := d.Convert2Entity()

		assert.NoError(t, err)

		err = svc.Store(dmn)

		assert.NoError(t, err)

	})
}

func TestOcurrencyRepository_FindAll(t *testing.T) {

	if err := godotenv.Load("../../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	database.Init()

	svc := OcurrencyRepository{}

	t.Run("Shouldn't return empty or error", func(t *testing.T) {
		dmn, err := svc.FindAll()
		assert.NotEmpty(t, dmn)
		assert.NoError(t, err)

	})
}

func TestOcurrencyRepository_FindByID(t *testing.T) {

	if err := godotenv.Load("../../.env"); err != nil {
		t.Error("Error loading .env file")
	}

	database.Init()

	svc := OcurrencyRepository{}

	t.Run("Shouldn't return an error", func(t *testing.T) {

		dmn, err := svc.FindByID(correct_id)
		assert.NotEmpty(t, dmn)
		assert.NoError(t, err)
	})

	t.Run("Shouldn't return an error", func(t *testing.T) {

		dmn, err := svc.FindByID(wrong_id)
		assert.Empty(t, dmn)
		assert.Error(t, err)
	})

}
