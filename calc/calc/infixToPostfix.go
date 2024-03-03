package calc

import (
	"errors"
	"unicode"
)

func isNumeric(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

var ops = make(map[string]int)

func InfixToPostfix(tokenList []string) (Stack, error) {
	var postfixList Stack = make([]string, 0)

	ops["+"] = 1
	ops["-"] = 1
	ops["/"] = 2
	ops["*"] = 2

	var opStack Stack = make([]string, 0)

	var tmp string

	for _, token := range tokenList {

		if isNumeric(token) {

			postfixList = postfixList.Push(token)

		} else if token == "(" {

			opStack = opStack.Push(token)

		} else if token == ")" {

			if len(opStack) == 0 {
				return nil, errors.New("не удалось спарсить выражение")
			}

			opStack, tmp = opStack.Pop()
			for tmp != "(" {
				postfixList = postfixList.Push(tmp)
				opStack, tmp = opStack.Pop()
			}

		} else {

			for len(opStack) > 0 && ops[opStack.Peek()] >= ops[token] {
				opStack, tmp = opStack.Pop()
				postfixList = postfixList.Push(tmp)
			}

			opStack = opStack.Push(token)
		}
	}

	for len(opStack) > 0 {
		opStack, tmp = opStack.Pop()
		postfixList = postfixList.Push(tmp)
	}

	return postfixList, nil

}
