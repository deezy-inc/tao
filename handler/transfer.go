package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) CreatePost(c echo.Context) (err error) {

	return c.Redirect(http.StatusMovedPermanently, "/post/")
}

func (h *Handler) FetchPost(c echo.Context) (err error) {

	return c.Render(http.StatusCreated, "post.html", map[string]interface{}{})
}

func (h *Handler) CheckPost(c echo.Context) (err error) {

	return c.Redirect(http.StatusMovedPermanently, "/post/")
}
