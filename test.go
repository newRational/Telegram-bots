package main

import (
	"articles-tbot/lib/e"
	"fmt"
	"github.com/ledongthuc/pdf"
	"log"
	"path/filepath"
)

func main() {
	content, err := readPdf(filepath.Join("schedules_storage", "Б21-505", "schedule.pdf"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
	return
}

func readPdf(path string) (string, error) {
	file, reader, err := pdf.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()
	var contentAsBytes []byte

	if _, err := reader.GetPlainText(); err != nil {
		return "", e.Wrap("can't read pdf file", err)
	}

	file.Close()
	return string(contentAsBytes), nil
}