package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Options struct {
	numOfAppearance bool
	repeated        bool
	notRepeated     bool
	ignoreCase      bool
	numFields       int
	numChars        int
}

func (options *Options) optionsInit() error {
	var err error
	flag.BoolVar(&options.numOfAppearance, "c", false, "number of line appearance")
	flag.BoolVar(&options.repeated, "d", false, "print only repeated lines")
	flag.BoolVar(&options.notRepeated, "u", false, "print only not repeated lines")
	flag.BoolVar(&options.ignoreCase, "i", false, "ignore letter case")
	flag.IntVar(&options.numFields, "f", 0, "skip first n fields in line")
	flag.IntVar(&options.numChars, "s", 0, "skip first n characters in line")
	flag.Parse()
	if options.numOfAppearance && options.repeated ||
		options.numOfAppearance && options.notRepeated ||
		options.repeated && options.notRepeated {
		err = errors.New("flags c, d, u can`t be together")
	}
	return err
}

// input
func setInputFile(inputAddress string) (*os.File, error) {
	inputAddress = flag.Args()[0]
	file, err := os.Open(inputAddress)
	return file, err
}

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
		file, err := setInputFile(inputAddress)
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

// uniq
func skipFields(text string, numFields int) string {
	lineArr := strings.Split(text, " ")
	copy(lineArr[0:], lineArr[numFields:])
	lineArr = lineArr[:len(lineArr)-numFields]
	return strings.Join(lineArr, " ")
}

func skipChars(text string, numChars int) string {
	if len(text) < numChars {
		return ""
	}
	charArr := strings.Split(text, "")
	copy(charArr[0:], charArr[numChars:])
	charArr = charArr[:len(charArr)-numChars]
	return strings.Join(charArr, "")
}

func wordHandler(text string, options Options) string {
	if options.ignoreCase {
		text = strings.ToLower(text)
	}
	if options.numFields > 0 {
		text = skipFields(text, options.numFields)
	}
	if options.numChars > 0 {
		text = skipChars(text, options.numChars)
	}
	return text
}

func uniq(text []string, options Options) ([]string, []int) {
	var repeatArr []int
	repeat := 0
	i := 0
	for ; i < len(text)-1; i++ {
		if wordHandler(text[i], options) == wordHandler(text[i+1], options) {
			repeat++
		} else {
			repeatArr = append(repeatArr, repeat+1)
			text = append(text[:i-repeat+1], text[i+1:]...)
			i -= repeat
			repeat = 0
		}
	}
	repeatArr = append(repeatArr, repeat+1)
	text = append(text[:i-repeat+1], text[i+1:]...)
	return text, repeatArr
}

// output
func write(ostream io.Writer, line string) error {
	_, err := io.WriteString(ostream, line)
	_, err = io.WriteString(ostream, "\n")
	return err
}

func setOutputFile(outputAddress string) (*os.File, error) {
	outputAddress = flag.Args()[1]
	file, err := os.Create(outputAddress)
	return file, err
}

func withRepeat(text string, repeat int) string {
	return strconv.Itoa(repeat) + " " + text
}

func output(text []string, repeatArr []int, optons Options) error {

	var err error
	outputAddress := ""
	var ostream io.Writer
	if flag.NArg() > 1 {
		file, err := setOutputFile(outputAddress)
		defer file.Close()
		if err != nil {
			return err
		}
		ostream = file
	} else {
		ostream = os.Stdout
	}

	for i, line := range text {
		switch {
		case optons.numOfAppearance:
			err = write(ostream, withRepeat(line, repeatArr[i]))
		case optons.repeated && repeatArr[i] > 1:
			err = write(ostream, line)
		case optons.notRepeated && repeatArr[i] == 1:
			err = write(ostream, line)
		case !optons.repeated && !optons.notRepeated:
			err = write(ostream, line)
		}
	}

	return err
}

// main
func main() {
	var options Options
	err := options.optionsInit()
	if err != nil {
		fmt.Println(err)
		return
	}

	text, err := input()
	if err != nil {
		fmt.Println(err)
		return
	}

	text, repeatArr := uniq(text, options)

	err = output(text, repeatArr, options)
	if err != nil {
		fmt.Println(err)
		return
	}
}
