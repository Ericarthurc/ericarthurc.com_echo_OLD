package utils

import "strings"

type Meta struct {
	FileName string
	Title    string
	Date     string
	Tags     []string
	Series   string
}

func newMeta(fileData *rawContent) *Meta {
	var meta Meta
	meta.FileName = fileData.fileName

	for _, v := range strings.Split(fileData.rawMeta, "\n") {
		switch strings.Split(v, ":")[0] {
		case "title":
			meta.Title = strings.TrimSpace(strings.Split(v, ":")[1])
		case "date":
			meta.Date = strings.TrimSpace(strings.Split(v, ":")[1])
		case "tags":
			for _, v := range strings.Split(strings.Split(v, ":")[1], ",") {
				meta.Tags = append(meta.Tags, strings.TrimSpace(v))
			}
		case "series":
			meta.Series = strings.TrimSpace(strings.Split(v, ":")[1])
		}
	}

	return &meta
}
