package api

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"github.com/sdil/learning/go/url-shortener/shortener"
)

type RedirectHandler interface {
	Get(echo.Context) error
	Post(echo.Context) error
}

type handler struct {
	redirectService shortener.RedirectService
}

func NewHandler(redirectService shortener.RedirectService) RedirectHandler {
	return &handler{redirectService: redirectService}
}

func (h *handler) Get(c echo.Context) error {
	code := c.Param("code")
	redirect, err := h.redirectService.Find(code)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, redirect.URL)
}

func (h *handler) Post(c echo.Context) error {
	redirect := &shortener.Redirect{}
	redirect.URL = c.FormValue("url")
	err := h.redirectService.Store(redirect)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, redirect.Code)
}