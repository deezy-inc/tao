package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/deezy-inc/tao/configs"
	"github.com/deezy-inc/tao/model"
	"github.com/deezy-inc/tao/tap"
	"github.com/labstack/echo"
)

type (
	Handler struct {
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func (h *Handler) Index(c echo.Context) (err error) {
	var (
		ctx = context.Background()
	)

	ctx, _ = configs.LoadConfig(ctx)
	tapClient := tap.NewClient(ctx)

	resp, err := tapClient.ListAssets() // leave 500000 cushion
	if err != nil {
		log.Println("Error opening channel")
		log.Println(err)
	}

	for i, a := range resp.TapAssets {
		resp.TapAssets[i] = decodeAssetFields(a)
	}

	// Retrieve featured (first 4) users from database
	users := []*model.PublicUser{}

	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"users":  users,
		"assets": resp.TapAssets,
	})
}
