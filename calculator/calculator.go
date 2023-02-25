package calculator

import (
	"bufio"
	"errors"
	"hw_1/Stack"
	"os"
	"strconv"
)

const SHIFT = 48

func makePriority() map[rune]int {
	return map[rune]int{
		'(': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'~': 3,
	}
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

func Input() (string, error) {
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
		if text[i] == '(' && text[i+1] == ')' {
			err = errors.New("error: brackets can`t be empty")
		}
		if text[i] == ')' && text[i+1] == '(' {
			err = errors.New("error: there is no operator between brackets")
		}
	}

	return text, err
}

func Parse(text string) string {
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
				stack.Push(rune('~'))
			} else {
				for stack.Len() > 0 && (operatorPriority[stack.Peek().(rune)] >= operatorPriority[rune(text[i])]) {
					postfix += string(stack.Pop().(rune)) + " "
				}
				stack.Push(rune(text[i]))
			}
		} else if text[i] == '(' {
			stack.Push(rune(text[i]))
		} else if text[i] == ')' {
			for stack.Len() > 0 && stack.Peek() != '(' {
				postfix += string(stack.Pop().(rune)) + " "
			}
			stack.Pop()
		}
	}
	for stack.Len() > 0 {
		postfix += string(stack.Pop().(rune)) + " "
	}

	return postfix
}

func Calculate(text string) (float64, error) {
	var err error
	stack := Stack.New()
	for i := 0; i < len(text); i++ {
		if isDigit(text[i]) {
			var str string
			for ; i < len(text) && isDigit(text[i]); i++ {
				str += string(text[i])
			}
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return 0.0, err
			}
			stack.Push(float64(num))
		} else if isOperator(text[i]) || text[i] == '~' {
			if text[i] == '~' {
				last := stack.Pop().(float64)
				stack.Push(float64(0 - last))
			} else {
				var first float64
				var second float64
				if stack.Len() != 0 {
					second = stack.Pop().(float64)
					first = stack.Pop().(float64)
				}
				switch text[i] {
				case '+':
					stack.Push(float64(first + second))
				case '-':
					stack.Push(float64(first - second))
				case '*':
					stack.Push(float64(first * second))
				case '/':
					if second == 0.0 {
						err = errors.New("error: division by zero")
						return 0.0, err
					}
					stack.Push(float64(first / second))
				}
			}
		}
	}
	return stack.Pop().(float64), err
}
