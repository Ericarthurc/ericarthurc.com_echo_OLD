package utils

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var markdown = goldmark.New(
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
	goldmark.WithRendererOptions(
		html.WithHardWraps(),
	),
)

func newMarkdown(fileData *rawContent) string {
	var buf bytes.Buffer

	if err := markdown.Convert([]byte(fileData.rawMarkdown), &buf); err != nil {
		return ""
	}

	return buf.String()
}
