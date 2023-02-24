package main

import (
	"bufio"
	"errors"
	"fmt"
	"hw_1/Stack"
	"os"
)

const SHIFT = 48

func makePriority() map[byte]int {
	return map[byte]int{
		'(': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'~': 3,
	}
}

func input() (string, error) {
	var text string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = scanner.Text()
	err := scanner.Err()

	if len(text) == 0 {
		err = errors.New("error: empty string")
	}

	if (len(text) != 0 && isOperator(text[0])) && text[0] != '-' {
		err = errors.New("error: first character can`t be binary operator")
	}

	bracketsNum := 0
	for i := range text {
		if !isDigit(text[i]) && !isBracket(text[i]) && !isOperator(text[i]) {
			err = errors.New("error: invalid input")
		}
		if text[i] == '(' {
			bracketsNum++
		}
		if text[i] == ')' {
			bracketsNum--
		}
	}
	if bracketsNum != 0 {
		err = errors.New("error: number of opening brackets not equal to number of closening brakets")
	}

	for i := 0; i < len(text)-1; i++ {
		if isOperator(text[i]) && isOperator(text[i+1]) {
			err = errors.New("error: two or more operators can`t go one by one")
		}
	}

	return text, err
}

func isBracket(char byte) bool {
	if char == '(' || char == ')' {
		return true
	} else {
		return false
	}
}

func isOperator(char byte) bool {
	if char == '+' || char == '-' || char == '*' || char == '/' {
		return true
	} else {
		return false
	}
}

func isDigit(char byte) bool {
	for i := 0; i < 10; i++ {
		if int(char)-SHIFT == i {
			return true
		}
	}
	return false
}

func parse(text string) string {
	var postfix string
	stack := Stack.New()
	operatorPriority := makePriority()

	i := 0
	for ; i < len(text); i++ {
		if isDigit(text[i]) {
			for ; i < len(text) && isDigit(text[i]); i++ {
				postfix += string(text[i])
			}
			i--
			postfix += " "
		}
		if isOperator(text[i]) {
			if (i == 0 || text[i-1] == '(') && text[i] == '-' {
				stack.Push('~')
			} else {
				for stack.Len() > 0 && (operatorPriority[stack.Peek()] >= operatorPriority[text[i]]) {
					postfix += string(stack.Pop()) + " "
				}
				stack.Push(text[i])
			}
		} else if text[i] == '(' {
			stack.Push(text[i])
		} else if text[i] == ')' {
			for stack.Len() > 0 && stack.Peek() != '(' {
				postfix += string(stack.Pop()) + " "
			}
			stack.Pop()
		}
	}
	for stack.Len() > 0 {
		postfix += string(stack.Pop()) + " "
	}

	return postfix
}



func main() {
	text, err := input()
	if err != nil {
		fmt.Println(err)
		return
	}

	a := parse(text)
	fmt.Println(a)

}
