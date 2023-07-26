package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/deezy-inc/tao/configs"
	"github.com/deezy-inc/tao/tap"
	"github.com/labstack/echo"
)

func (h *Handler) Address(c echo.Context) (err error) {
	var (
		ctx = context.Background()
		a   = &tap.TapAddressRequest{}
	)

	ctx, _ = configs.LoadConfig(ctx)
	tapClient := tap.NewClient(ctx)

	c.Bind(a)

	resp, err := tapClient.CreateAddress(a.AssetId, a.Amount)
	if err != nil {
		log.Println("Error creating address")
		log.Println(err)
	}

	return c.Render(http.StatusOK, "address.html", map[string]interface{}{
		"address": resp,
	})
}

func (h *Handler) ViewAddress(c echo.Context) (err error) {
	return c.Render(http.StatusOK, "address.html", map[string]interface{}{})
}
