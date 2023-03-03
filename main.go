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
	file := os.Stdin
	if len(inputAddress) != 0 {
		file, err = os.Open(inputAddress)
		if err != nil {
			return text, err
		}
		defer file.Close()
	}
	text, err = scanText(file)
	return text, err
}

func write(ostream io.Writer, text []string) error {
	var err error
	for _, line := range text {
		_, err = io.WriteString(ostream, line)
		_, err = io.WriteString(ostream, "\n")
	}
	return err
}

func output(result []string, outputAddress string) error {
	var err error
	file := os.Stdout
	if len(outputAddress) != 0 {
		file, err = os.Create(outputAddress)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	err = write(file, result)
	return err
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

	result, err := uniq.Uniq(text, options)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = output(result, fileNames.OutputAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
}
