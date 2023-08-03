package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/deezy-inc/tao/configs"
	"github.com/deezy-inc/tao/tap"
	"github.com/labstack/echo"
)

func (h *Handler) Send(c echo.Context) (err error) {
	var (
		ctx = context.Background()
		a   = &tap.TapSendRequest{}
	)

	ctx, _ = configs.LoadConfig(ctx)
	tapClient := tap.NewClient(ctx)

	c.Bind(a)

	_, err = tapClient.CreateSend(a.Address)
	if err != nil {
		log.Println("Error creating address")
		log.Println(err)
	}

	return c.Render(http.StatusOK, "send.html", map[string]interface{}{
		"address": "Success!",
	})
}

func (h *Handler) ViewSend(c echo.Context) (err error) {
	return c.Render(http.StatusOK, "send.html", map[string]interface{}{})
}
