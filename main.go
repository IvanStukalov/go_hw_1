package main

import (
	"bufio"
	"fmt"
	"hw_1/uniq"
	"io"
	"os"
)

// input
func scanText(istream io.Reader) ([]string, error) {
	var text []string
	scanner := bufio.NewScanner(istream)

	for i := 0; scanner.Scan(); i++ {
		text = append(text, scanner.Text())
	}
	err := scanner.Err()

	return text, err
}

func input(inputAddress string) ([]string, error) {
	var err error
	var text []string
	var istream io.Reader
	if len(inputAddress) != 0 {
		file, err := os.Open(inputAddress)
		if err != nil {
			return text, err
		}
		defer file.Close()
		istream = file
	} else {
		istream = os.Stdin
	}
	text, err = scanText(istream)
	return text, err
}

// main
func main() {
	var options uniq.Options
	var fileNames uniq.FileNames
	err := uniq.OptionsInit(&options, &fileNames)
	if err != nil {
		fmt.Println(err)
		return
	}
	text, err := input(fileNames.InputAddress)
	if err != nil {
		fmt.Println(err)
		return
	}

	text, repeatArr := uniq.Uniq(text, options)

	err = uniq.Output(text, repeatArr, options, fileNames.OutputAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
}
