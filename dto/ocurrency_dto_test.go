package dto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOcurrencyDTO_Convert2Entity(t *testing.T) {

	t.Run("Shouldn't return an error", func(t *testing.T) {
		dto := OcurrencyDTO{
			Amount: 5,
			From:   "BRL",
			To:     "USD",
			Rate:   4.5,
		}

		dmn, err := dto.Convert2Entity()
		assert.NotEmpty(t, dmn)
		assert.NoError(t, err)
	})

	t.Run("Should return an error because ocurrency is invalid", func(t *testing.T) {
		dto := OcurrencyDTO{
			Amount: 0,
			From:   "LTC",
			To:     "USD",
			Rate:   4.5,
		}

		_, err := dto.Convert2Entity()
		assert.Error(t, err)
	})

}
