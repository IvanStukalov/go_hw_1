package main

import (
	"bufio"
	"fmt"
	"hw_1/calculator"
	"os"
)

func Input() (string, error) {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	err := scanner.Err()
	if err != nil {
		return text, err
	}
	return text, err
}

func main() {
	text, err := Input()
	if err != nil {
		fmt.Println(err)
		return
	}
	parsed, err := calculator.Parse(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := calculator.Calculate(parsed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
