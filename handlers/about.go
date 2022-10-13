package handlers

import (
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
)

func AboutIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "about", pongo2.Context{"url": c.Request().URL.Path})
}
