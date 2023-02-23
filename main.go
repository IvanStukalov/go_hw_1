package main

import (
	"bufio"
	"errors"
	"flag"
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

func input() ([]string, error) {
	var err error
	var text []string
	if flag.NArg() > 2 {
		err = errors.New("too many arguments")
	}
	inputAddress := ""
	var istream io.Reader
	if flag.NArg() > 0 {
		inputAddress = flag.Args()[0]
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
	err := uniq.OptionsInit(&options)
	if err != nil {
		fmt.Println(err)
		return
	}
	text, err := input()
	if err != nil {
		fmt.Println(err)
		return
	}

	text, repeatArr := uniq.Uniq(text, options)

	err = uniq.Output(text, repeatArr, options)
	if err != nil {
		fmt.Println(err)
		return
	}
}
