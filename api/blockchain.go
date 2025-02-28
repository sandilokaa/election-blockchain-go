package api

import (
	"election-blockchain-go/domain"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type blockchainApi struct {
	bc *domain.Blockchain
}

func NewBlockchain(app *fiber.App, bc *domain.Blockchain) {
	bca := blockchainApi{
		bc: bc,
	}

	app.Get("/chain", bca.Chain)
	app.Post("/give-mandate", bca.GiveMandate)
	app.Get("/check-mandate", bca.CheckMandate)
}

func (bca blockchainApi) Chain(ctx *fiber.Ctx) error {
	return ctx.JSON(Response[[]*domain.Block]{Message: "success", Data: bca.bc.Chain})
}

func (bca blockchainApi) GiveMandate(ctx *fiber.Ctx) error {
	var req domain.Mandate
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	isSuccess := bca.bc.GiveMandate(req.From, req.To, req.Value)
	if isSuccess {
		return ctx.JSON(Response[string]{Message: "success give mandate"})
	}

	return ctx.Status(http.StatusBadRequest).JSON(Response[string]{Message: "insufficient mandate"})
}

func (bca blockchainApi) CheckMandate(ctx *fiber.Ctx) error {
	q := ctx.Query("q")

	data := make(map[string]int64)
	for _, v := range strings.Split(q, ",") {
		data[v] = bca.bc.CalculateMandate(v)
	}

	return ctx.JSON(Response[map[string]int64]{Message: "success", Data: data})
}
