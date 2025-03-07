package utils

import (
	"bufio"
	"compress/gzip"
	"os"
)

type document struct {
	Title string `xml:"title"`
	ID    int
}

func LoadDocuments(path string) ([]document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	var docs []document
	scanner := bufio.NewScanner(gz)
	id := 0
	for scanner.Scan() {
		title := scanner.Text()
		if title != "" {
			docs = append(docs, document{
				Title: title,
				ID:    id,
			})
			id++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return docs, nil
}
