package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware() fiber.Handler {
	err := cors.New(cors.Config{
		Next: func(c *fiber.Ctx) bool {
			if string(c.Request().Header.Method()) == "OPTIONS" {
				c.Status(fiber.StatusNoContent)
				return false
			}
			return true
		},
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, HEAD",
		AllowHeaders:     "*",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
	return err
}
