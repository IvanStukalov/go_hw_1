package main

import (
	"bufio"
	"fmt"
	"os"
)

func input() (string, error) {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	err := scanner.Err()

	return text, err
}

func main() {
	text, err := input()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
}
