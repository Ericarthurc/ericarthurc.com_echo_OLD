package router

import (
	"ericarthurc/handlers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	blog := e.Group("/")
	blog.GET("", handlers.BlogIndex)
	blog.GET("blog/:blog", handlers.GetBlog)

	series := e.Group("/series")
	series.GET("", handlers.SeriesIndex)
	series.GET("/:series", handlers.GetSeries)

	category := e.Group("/category")
	category.GET("/:category", handlers.GetCategory)

	about := e.Group("/about")
	about.GET("", handlers.AboutIndex)

	// admin := e.Group("/admin")
	// admin.GET("/")
}
