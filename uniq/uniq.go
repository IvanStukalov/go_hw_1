package uniq

import (
	"errors"
	"flag"
	"io"
	"os"
	"strconv"
	"strings"
)

type Options struct {
	NumOfAppearance bool
	Repeated        bool
	NotRepeated     bool
	IgnoreCase      bool
	NumFields       int
	NumChars        int
}

type FileNames struct {
	InputAddress  string
	OutputAddress string
}

func OptionsInit(options *Options, fileNames *FileNames) error {
	var err error
	flag.BoolVar(&options.NumOfAppearance, "c", false, "number of line appearance")
	flag.BoolVar(&options.Repeated, "d", false, "print only repeated lines")
	flag.BoolVar(&options.NotRepeated, "u", false, "print only not repeated lines")
	flag.BoolVar(&options.IgnoreCase, "i", false, "ignore letter case")
	flag.IntVar(&options.NumFields, "f", 0, "skip first n fields in line")
	flag.IntVar(&options.NumChars, "s", 0, "skip first n characters in line")
	flag.Parse()
	if options.NumOfAppearance && options.Repeated ||
		options.NumOfAppearance && options.NotRepeated ||
		options.Repeated && options.NotRepeated {
		err = errors.New("flags c, d, u can`t be together")
		flag.Usage()
	}

	if flag.NArg() > 2 {
		err = errors.New("too many arguments")
		return err
	}
	if flag.NArg() > 0 {
		fileNames.InputAddress = flag.Args()[0]
	}
	if flag.NArg() > 1 {
		fileNames.OutputAddress = flag.Args()[1]
	}

	return err
}

// uniq
func skipFields(text string, numFields int) string {
	lineArr := strings.Fields(text)
	if len(lineArr) <= numFields {
		numFields = len(lineArr)
	}
	lineArr = lineArr[numFields:]
	return strings.Join(lineArr, " ")
}

func skipChars(text string, numChars int) string {
	if len(text) < numChars {
		return ""
	}
	return text[numChars:]
}

func wordHandler(text string, options Options) string {
	if options.IgnoreCase {
		text = strings.ToLower(text)
	}
	if options.NumFields > 0 {
		text = skipFields(text, options.NumFields)
	}
	if options.NumChars > 0 {
		text = skipChars(text, options.NumChars)
	}
	return text
}

func Uniq(text []string, options Options) ([]string, []int) {
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

func withRepeat(text string, repeat int) string {
	return strconv.Itoa(repeat) + " " + text
}

func Output(text []string, repeatArr []int, optons Options, outputAddress string) error {
	var err error
	var ostream io.Writer
	if len(outputAddress) != 0 {
		file, err := os.Create(outputAddress)
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
		case optons.NumOfAppearance:
			err = write(ostream, withRepeat(line, repeatArr[i]))
		case optons.Repeated && repeatArr[i] > 1:
			err = write(ostream, line)
		case optons.NotRepeated && repeatArr[i] == 1:
			err = write(ostream, line)
		case !optons.Repeated && !optons.NotRepeated:
			err = write(ostream, line)
		}
	}

	return err
}
