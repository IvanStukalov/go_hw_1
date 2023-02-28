package uniq

import (
	"errors"
	"flag"
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

func Uniq(text []string, options Options) ([]string, error) {
	var repeatArr []int
	var result []string
	var err error

	if options.NumOfAppearance && options.Repeated ||
		options.NumOfAppearance && options.NotRepeated ||
		options.Repeated && options.NotRepeated {
		err = errors.New("flags c, d, u can`t be together")
		return result, err
	}

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

	for i, line := range text {
		switch {
		case options.NumOfAppearance:
			result = append(result, withRepeat(line, repeatArr[i]))
		case options.Repeated:
			if repeatArr[i] > 1 {
				result = append(result, line)
			}
		case options.NotRepeated:
			if repeatArr[i] == 1 {
				result = append(result, line)
			}
		default:
			result = append(result, line)
		}
	}
	return result, err
}

func withRepeat(text string, repeat int) string {
	return strconv.Itoa(repeat) + " " + text
}
