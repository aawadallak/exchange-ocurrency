package controllers

import (
	"exchange_currency/dto"
	"exchange_currency/infra/repository"
	"exchange_currency/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type OcurrencyController struct{}

func (o *OcurrencyController) Exchange(c *fiber.Ctx) error {

	amount, err := strconv.ParseFloat(c.Params("amount"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "parãmetro inválido para quantidade, por favor, tente novamente"})
	}

	rate, err := strconv.ParseFloat(c.Params("rate"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "parãmetro inválido para o cambio, por favor, tente novamente"})
	}

	d := dto.OcurrencyDTO{
		Amount: amount,
		From:   c.Params("from"),
		To:     c.Params("to"),
		Rate:   rate,
	}

	svc := usecases.NewOcurrencyService(&repository.OcurrencyRepository{})

	r, err := svc.Exchange(&d)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(r)
}

func (o *OcurrencyController) FindAllOcurrencies(c *fiber.Ctx) error {
	svc := usecases.NewOcurrencyService(&repository.OcurrencyRepository{})

	r, err := svc.FindAll()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if r == nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
	}

	return c.Status(fiber.StatusOK).JSON(r)
}

func (o *OcurrencyController) FindOcurrencyByID(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.Params("id"), 10, 10)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "id inválido, tente novamente!"})
	}

	svc := usecases.NewOcurrencyService(&repository.OcurrencyRepository{})

	r, err := svc.FindByID(uint(id))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(r)
}
