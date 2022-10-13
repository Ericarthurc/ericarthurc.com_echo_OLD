package handlers

import (
	"ericarthurc/utils"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
)

func SeriesIndex(c echo.Context) error {
	series, err := utils.GetSeriesSlice()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.Render(http.StatusOK, "index_series", pongo2.Context{"seriesList": series, "url": c.Request().URL.Path})
}

func GetSeries(c echo.Context) error {
	series := c.Param("series")

	seriesList, err := utils.GetSliceBySeriesParam(series)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.Render(http.StatusOK, "series", pongo2.Context{"series": series, "seriesList": seriesList, "url": c.Request().URL.Path})
}
