package calculator

import (
	"errors"
	"fmt"
	"hw_1/stack"
	"strconv"
	"unicode"
)

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
	return char == '(' || char == ')'
}

func isOperator(char byte) bool {
	return char == '+' || char == '-' || char == '*' || char == '/'
}

func isDot(char byte) bool {
	return char == '.'
}

func validate(text string) error {
	if len(text) == 0 {
		return fmt.Errorf("error: empty string")
	}
	if (len(text) != 0 && isOperator(text[0])) && text[0] != '-' {
		return fmt.Errorf("error: first character can`t be binary operator")
	}
	if len(text) != 0 && isDot(text[0]) {
		return fmt.Errorf("error: first character can`t be dot")
	}

	bracketsNum := 0
	for i := range text {
		if !unicode.IsDigit(rune(text[i])) && !isBracket(text[i]) && !isOperator(text[i]) && !isDot(text[i]) {
			return fmt.Errorf("error: invalid input")
		}
		if text[i] == '(' {
			bracketsNum++
		}
		if text[i] == ')' {
			bracketsNum--
		}
	}
	if bracketsNum != 0 {
		return fmt.Errorf("error: number of opening brackets not equal to number of closening brakets")
	}

	for i := 0; i < len(text)-1; i++ {
		if isOperator(text[i]) && isOperator(text[i+1]) {
			return fmt.Errorf("error: two or more operators can`t go one by one")
		}
		if isDot(text[i]) && isDot(text[i+1]) {
			return fmt.Errorf("error: two or more dots can`t go one by one")
		}
		if isDot(text[i]) && (isOperator(text[i+1]) || isBracket(text[i+1])) {
			return fmt.Errorf("error: operator can`t go after dot")
		}
		if (isOperator(text[i]) || isBracket(text[i])) && isDot(text[i+1]) {
			return fmt.Errorf("error: dot can`t go after operator")
		}
		if text[i] == '(' && text[i+1] == ')' {
			return fmt.Errorf("error: brackets can`t be empty")
		}
		if text[i] == ')' && text[i+1] == '(' {
			return fmt.Errorf("error: there is no operator between brackets")
		}
	}
	return nil
}

func Parse(text string) (string, error) {
	err := validate(text)
	if err != nil {
		return text, err
	}

	var postfix string
	stack := stack.New()
	operatorPriority := makePriority()

	i := 0
	for ; i < len(text); i++ {
		if unicode.IsDigit(rune(text[i])) {
			for ; i < len(text) && (unicode.IsDigit(rune(text[i])) || isDot(text[i])); i++ {
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

	return postfix, err
}

func Calculate(text string) (float64, error) {
	var err error
	stack := stack.New()
	for i := 0; i < len(text); i++ {
		if unicode.IsDigit(rune(text[i])) {
			var str string
			for ; i < len(text) && (unicode.IsDigit(rune(text[i])) || isDot(text[i])); i++ {
				str += string(text[i])
			}
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				return 0.0, err
			}
			stack.Push(num)
		} else if isOperator(text[i]) || text[i] == '~' {
			if text[i] == '~' {

				last, ok := stack.Pop().(float64)
				if !ok {
					return 0.0, fmt.Errorf("error: number is not float64")
				}

				stack.Push(0.0 - last)
			} else {

				second, ok := stack.Pop().(float64)
				if !ok {
					return 0.0, fmt.Errorf("error: number is not float64")
				}

				first, ok := stack.Pop().(float64)
				if !ok {
					return 0.0, fmt.Errorf("error: number is not float64")
				}

				switch text[i] {
				case '+':
					stack.Push(first + second)
				case '-':
					stack.Push(first - second)
				case '*':
					stack.Push(first * second)
				case '/':
					if second == 0.0 {
						err = errors.New("error: division by zero")
						return 0.0, err
					}
					stack.Push(first / second)
				}
			}
		}
	}
	res, ok := stack.Pop().(float64)
	if !ok {
		return 0.0, fmt.Errorf("error: number is not float64")
	}
	return res, err
}
