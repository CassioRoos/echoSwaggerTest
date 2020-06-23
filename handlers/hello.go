package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Hello godoc
// @Description Hello World
// @Summary Say Hello to the world :D.
// @Success 200 {string} string "ok"
// @Router /hello [get]
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
