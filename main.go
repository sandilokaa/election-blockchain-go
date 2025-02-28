package main

import (
	"election-blockchain-go/api"
	"election-blockchain-go/domain"

	"github.com/gofiber/fiber/v2"
)

func main() {
	bc := domain.NewBlockchain()

	go bc.PlenaryRecap()

	app := fiber.New()
	api.NewBlockchain(app, bc)

	_ = app.Listen(":8000")
}
