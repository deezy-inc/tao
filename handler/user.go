package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func (h *Handler) Signup(c echo.Context) (err error) {

	return c.Redirect(http.StatusMovedPermanently, "/dashboard")
}

func (h *Handler) ViewSignup(c echo.Context) (err error) {
	return c.Render(http.StatusOK, "signup.html", map[string]interface{}{})
}

func (h *Handler) Login(c echo.Context) (err error) {

	return c.Redirect(http.StatusMovedPermanently, "/dashboard")
}

func (h *Handler) ViewLogin(c echo.Context) (err error) {
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{})
}

func (h *Handler) ListUsers(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) FetchUser(c echo.Context) (err error) {

	return c.Render(http.StatusOK, "user.html", map[string]interface{}{
		"user": nil,
	})
}

func (h *Handler) Dashboard(c echo.Context) (err error) {

	return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{})
}

func (h *Handler) UpdateUser(c echo.Context) (err error) {

	return c.Redirect(http.StatusMovedPermanently, "/dashboard")
}

func (h *Handler) CreateWithdrawal(c echo.Context) (err error) {

	return c.Redirect(http.StatusMovedPermanently, "/dashboard")
}
