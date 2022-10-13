package utils

import (
	"os"
	"sort"
	"strings"
	"time"
)

func GetBlogPostByName(fileName string) (markdown string, meta *Meta, err error) {
	rawData, err := newRawContent(fileName)
	if err != nil {
		return "", nil, err
	}

	return newMarkdown(rawData), newMeta(rawData), nil
}

func GetBlogPostsSlice() ([]*Meta, error) {
	var blogMetaSlice []*Meta

	files, err := os.ReadDir("./blogPosts")
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		rawData, err := newRawContent(strings.TrimSuffix(f.Name(), ".md"))
		if err != nil {
			return nil, err
		}

		meta := newMeta(rawData)

		blogMetaSlice = append(blogMetaSlice, meta)
	}

	// Sort by date and alphabetical
	sort.Slice(blogMetaSlice, func(i, j int) bool {
		format := "January 2, 2006"
		t1, _ := time.Parse(format, blogMetaSlice[i].Date)
		t2, _ := time.Parse(format, blogMetaSlice[j].Date)

		if t1.Unix() == t2.Unix() {
			return blogMetaSlice[i].Title < blogMetaSlice[j].Title
		} else {
			return t1.Unix() > t2.Unix()
		}
	})

	return blogMetaSlice, nil
}

func GetSeriesSlice() ([]string, error) {
	var seriesSlice []string

	files, err := os.ReadDir("./blogPosts")
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		rawData, err := newRawContent(strings.TrimSuffix(f.Name(), ".md"))
		if err != nil {
			return nil, err
		}

		meta := newMeta(rawData)

		// check if series is "" and skip
		// go uses "" as a zero value for the string type
		// the "" will get pushed to the slice if not checked
		if meta.Series == "" {
			continue
		}

		seriesSlice = append(seriesSlice, meta.Series)
	}

	// Remove duplicate strings in slice
	var uniqueSeriesSlice []string
	for _, s := range seriesSlice {
		var same bool
		for _, a := range uniqueSeriesSlice {
			if a == s {
				same = true
				continue
			}
		}

		if !same {
			uniqueSeriesSlice = append(uniqueSeriesSlice, s)
		}
	}

	// Sort strings in alphabetical order
	sort.Strings(uniqueSeriesSlice)

	return uniqueSeriesSlice, nil
}

func GetSliceBySeriesParam(series string) ([]*Meta, error) {
	var seriesMetaSlice []*Meta

	files, err := os.ReadDir("./blogPosts")
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		rawData, err := newRawContent(strings.TrimSuffix(f.Name(), ".md"))
		if err != nil {
			return nil, err
		}

		meta := newMeta(rawData)

		if meta.Series == series {
			seriesMetaSlice = append(seriesMetaSlice, meta)
		}
	}

	// Sort by date and alphabetical
	sort.Slice(seriesMetaSlice, func(i, j int) bool {
		format := "January 2, 2006"
		t1, _ := time.Parse(format, seriesMetaSlice[i].Date)
		t2, _ := time.Parse(format, seriesMetaSlice[j].Date)

		if t1.Unix() == t2.Unix() {
			return seriesMetaSlice[i].Title < seriesMetaSlice[j].Title
		} else {
			return t1.Unix() > t2.Unix()
		}
	})

	return seriesMetaSlice, nil
}

func GetSliceByCategoryParam(category string) ([]*Meta, error) {
	var categoryMetaSlice []*Meta

	files, err := os.ReadDir("./blogPosts")
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		rawData, err := newRawContent(strings.TrimSuffix(f.Name(), ".md"))
		if err != nil {
			return nil, err
		}

		meta := newMeta(rawData)

		for _, t := range meta.Tags {
			if t == category {
				categoryMetaSlice = append(categoryMetaSlice, meta)
			}
		}
	}

	// Sort by date and alphabetical
	sort.Slice(categoryMetaSlice, func(i, j int) bool {
		format := "January 2, 2006"
		t1, _ := time.Parse(format, categoryMetaSlice[i].Date)
		t2, _ := time.Parse(format, categoryMetaSlice[j].Date)

		if t1.Unix() == t2.Unix() {
			return categoryMetaSlice[i].Title < categoryMetaSlice[j].Title
		} else {
			return t1.Unix() > t2.Unix()
		}
	})

	return categoryMetaSlice, nil
}
