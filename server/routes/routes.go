package routes

import (
	"exchange_currency/app/http/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetRoute(router *fiber.App) *fiber.App {

	controller := controllers.OcurrencyController{}

	router.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Pong"})
	})

	exchange := router.Group("/exchange")
	{
		exchange.Get("/:amount/:from/:to/:rate", controller.Exchange)

		exchange.Get("/all", controller.FindAllOcurrencies)

		exchange.Get("/:id", controller.FindOcurrencyByID)
	}

	return router
}
