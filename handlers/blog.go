package handlers

import (
	"ericarthurc/utils"
	"fmt"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
)

func BlogIndex(c echo.Context) error {
	posts, err := utils.GetBlogPostsSlice()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.Render(http.StatusOK, "index", pongo2.Context{"blogList": posts, "url": c.Request().URL.Path})
}

func GetBlog(c echo.Context) error {
	blogName := c.Param("blog")

	markdown, meta, err := utils.GetBlogPostByName(blogName)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("blog post %s was not found", blogName))
	}

	return c.Render(http.StatusOK, "blog", pongo2.Context{"markdown": markdown, "meta": meta, "url": c.Request().URL.Path})
}
