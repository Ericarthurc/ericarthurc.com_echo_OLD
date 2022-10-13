package utils

import (
	"io"
	"path"

	"github.com/flosch/pongo2/v6"
	"github.com/labstack/echo/v4"
)

type RenderOptions struct {
	TemplateDir       string
	TemplateExtension string
}

type Template struct {
	Options *RenderOptions
}

func NewPongo(options RenderOptions) *Template {
	return &Template{
		Options: &options,
	}
}

func (t *Template) Render(w io.Writer, name string, data any, c echo.Context) error {
	var template *pongo2.Template
	filename := path.Join(t.Options.TemplateDir, name)
	template = pongo2.Must(pongo2.FromCache(filename + "." + t.Options.TemplateExtension))

	return template.ExecuteWriter(data.(pongo2.Context), w)
}
