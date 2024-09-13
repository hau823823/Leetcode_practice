package evaluateReversePolishNotation

import (
	"strconv"
)

func EvalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, num)
		} else {
			num1 := stack[len(stack)-2]
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			if token == "+" {
				stack = append(stack, num1+num2)
			} else if token == "-" {
				stack = append(stack, num1-num2)
			} else if token == "*" {
				stack = append(stack, num1*num2)
			} else if token == "/" {
				stack = append(stack, num1/num2)
			}
		}
	}

	return stack[0]
}
