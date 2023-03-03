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

func OptionsInit(options *uniq.Options, fileNames *uniq.FileNames) error {
	var err error
	flag.BoolVar(&options.NumOfAppearance, "c", false, "number of line appearance")
	flag.BoolVar(&options.Repeated, "d", false, "print only repeated lines")
	flag.BoolVar(&options.NotRepeated, "u", false, "print only unique lines")
	flag.BoolVar(&options.IgnoreCase, "i", false, "ignore letter case")
	flag.IntVar(&options.NumFields, "f", 0, "skip first n fields in line")
	flag.IntVar(&options.NumChars, "s", 0, "skip first n characters in line")
	flag.Parse()

	if flag.NArg() > 2 {
		err = errors.New("too many arguments")
		return err
	}
	fileNames.InputAddress = flag.Arg(0)
	fileNames.OutputAddress = flag.Arg(1)
	return err
}

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
	err := OptionsInit(&options, &fileNames)
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
