package main

import (
	"bufio"
	"log"
	"os"
)

type FileTickerStore struct{}

func (s FileTickerStore) Read() ([]string, error) {
	f := OpenTickerFile("tickers.txt")
	scanner := bufio.NewScanner(f)
	var tickers []string

	defer f.Close()

	for scanner.Scan() {
		tickers = append(tickers, scanner.Text())
	}

	return tickers, scanner.Err()
}

func OpenTickerFile(filename string) *os.File {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	return f
}
