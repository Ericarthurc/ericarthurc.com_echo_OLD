package utils

import (
	"os"
	"strings"
)

type rawContent struct {
	rawMeta     string
	rawMarkdown string
	fileName    string
}

func newRawContent(fileName string) (*rawContent, error) {
	raw, err := os.ReadFile("./blogPosts/" + fileName + ".md")
	if err != nil {
		return nil, err
	}

	// Split file contents into raw parts
	dividedRaw := strings.Split(string(raw), "---")

	return &rawContent{rawMeta: dividedRaw[1], rawMarkdown: dividedRaw[2], fileName: fileName}, nil
}
