package main

import (
	"github.com/gofiber/fiber/v2"
)

type Quote struct {
	Quote string
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		a := fiber.AcquireAgent()
		req := a.Request()
		req.Header.SetMethod(fiber.MethodGet)
		req.SetRequestURI("https://quotes-production.up.railway.app/api/quote")

		if err := a.Parse(); err != nil {
			panic(err)
		}

		var d Quote

		code, _, errs := a.Struct(&d)

		if code != 200 {
			panic(errs)
		}

		return c.JSON(d)
	})

	app.Listen("0.0.0.0:3001")
}
