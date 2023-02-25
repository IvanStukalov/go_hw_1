package main

import (
	"fmt"
	"hw_1/calculator"
)

func main() {
	text, err := calculator.Input()
	if err != nil {
		fmt.Println(err)
		return
	}
	parsed := calculator.Parse(text)
	fmt.Println(parsed)
	
	res, err := calculator.Calculate(parsed)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
