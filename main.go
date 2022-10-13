package main

import (
	"ericarthurc/router"
	"ericarthurc/utils"
	"fmt"
	"net/http"
	"os"

	"github.com/flosch/pongo2/v6"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Register environmental variables
	godotenv.Load()

	e := echo.New()

	// Register pongo view engine
	e.Renderer = utils.NewPongo(utils.RenderOptions{
		TemplateDir:       "views",
		TemplateExtension: "dj",
	})

	// Register custom error handler
	e.HTTPErrorHandler = customErrorHandler

	// Register static files
	e.Static("/public", "./public")
	e.File("/sw.js", "./public/js/sw.js")
	e.File("/robots.txt", "./public/robots.txt")

	// Register global middleware
	e.Use(middleware.Recover())

	// Register routes
	router.SetupRoutes(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func customErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	c.Logger().Error(err)
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}

	c.Render(200, "error", pongo2.Context{
		"code":  code,
		"path":  c.Request().URL,
		"error": err.Error(),
	})
}
