package handlers

import (
	"ericarthurc/utils"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
)

func GetCategory(c echo.Context) error {
	category := c.Param("category")

	categoryList, err := utils.GetSliceByCategoryParam(category)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.Render(http.StatusOK, "category", pongo2.Context{"category": category, "categoryList": categoryList, "url": c.Request().URL.Path})
}
