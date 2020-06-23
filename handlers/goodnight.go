package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"time"
)

type (
	GoodNight struct {
		Date time.Time `json:"date"`
		Text string    `json:"text"`
	}

	HTTPError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

// GetSayGoodNigth godoc
// @Description Say GoodNight to someone
// @Summary Be polite and say goodnight to a friend.
// @Param to path string true "The person whom you want to say goodnight"
// @Success 200 {object} GoodNight "GoodNight"
// @Failure 400 {object} HTTPError
// @Router /goodnight/{to} [get]
func GetSayGoodNigth(ctx echo.Context) error {
	text := ctx.Param("to")
	if strings.TrimSpace(text) == "" {
		return ctx.JSON(http.StatusBadRequest, HTTPError{http.StatusBadRequest, "You must define to whom you want to say goodbye"})
	}
	return ctx.JSON(http.StatusOK, &GoodNight{Date: time.Now(), Text: fmt.Sprintf("Good night %s", text)})
}
