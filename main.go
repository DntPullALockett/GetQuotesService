package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type Quote struct {
	Quote string
}

func main() {
	app := fiber.New()

	viper.SetConfigFile("ENV")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	port := fmt.Sprint(viper.Get("PORT"))

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

	app.Listen("0.0.0.0:" + port)
}
