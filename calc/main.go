package main

import (
	"flag"
	"fmt"
	"strconv"
)

import "unicode"

var result = make([]string, 0)
var numberBuffer = make([]rune, 0)
var operatorBuffer = make([]rune, 0)

func emptyNumberBufferAsLiteral() {
	if len(numberBuffer) > 0 {
		result = append(result, string(numberBuffer))
		numberBuffer = nil
	}
}

func emptyOperatorBuffer() {
	if len(operatorBuffer) > 0 {
		result = append(result, string(operatorBuffer))
		operatorBuffer = nil
	}
}

func isLeftParenthesis(ch rune) bool {
	return ch == '('
}

func isRightParenthesis(ch rune) bool {
	return ch == ')'
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func isOperator(ch rune) bool {
	if ch == '+' {
		return true
	}
	if ch == '-' {
		return true
	}
	if ch == '*' {
		return true
	}
	if ch == '/' {
		return true
	}

	return false
}

func isNumeric(str string) bool {
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func tokenize(raw string) {
	var str = []rune(raw)

	for _, ch := range str {

		if isDigit(ch) {

			numberBuffer = append(numberBuffer, ch)

			if len(operatorBuffer) > 0 {
				emptyOperatorBuffer()
			}

		} else if isOperator(ch) {

			emptyNumberBufferAsLiteral()
			operatorBuffer = append(operatorBuffer, ch)

		} else if isLeftParenthesis(ch) {

			if len(numberBuffer) > 0 {
				emptyNumberBufferAsLiteral()
				result = append(result, "*")
			}

			if len(operatorBuffer) > 0 {
				emptyOperatorBuffer()
			}

			result = append(result, "(")

		} else if isRightParenthesis(ch) {

			emptyNumberBufferAsLiteral()
			result = append(result, ")")

		}

	}
	if len(numberBuffer) > 0 {
		emptyNumberBufferAsLiteral()
	}

	if len(operatorBuffer) > 0 {
		emptyOperatorBuffer()
	}
}

type stack []string

func (s stack) Push(v string) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() string {
	l := len(s)
	return s[l-1]
}

var ops = make(map[string]int)

func infixToPostfix(tokenList []string) (postfixList stack) {
	ops["("] = 0
	ops["+"] = 1
	ops["-"] = 1
	ops["/"] = 2
	ops["*"] = 2

	var opStack stack = make([]string, 0)

	var tmp string

	for _, token := range tokenList {

		if isNumeric(token) {

			postfixList = postfixList.Push(token)

		} else if token == "(" {

			opStack = opStack.Push(token)

		} else if token == ")" {

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

	return postfixList

}

func ConvertToInt(x string) int {
	var i, _ = strconv.Atoi(x)
	return i
}

func evalBinary(x string, y string, op string) int {
	if op == "+" {
		return ConvertToInt(x) + ConvertToInt(y)
	}

	if op == "-" {
		return ConvertToInt(x) - ConvertToInt(y)
	}

	if op == "*" {
		return ConvertToInt(x) * ConvertToInt(y)
	}

	if op == "/" {
		return ConvertToInt(x) / ConvertToInt(y)
	}

	return -1
}

func eval(tokens []string) int {
	var res stack = make([]string, 0)

	var x, y string

	for _, token := range tokens {

		var _, ok = ops[token] // Если встретился оператор (+, -, *, /)

		if ok {

			res, y = res.Pop()
			res, x = res.Pop()

			res = res.Push(strconv.Itoa(evalBinary(x, y, token)))

		} else { // Если встретилось число

			res = res.Push(token)

		}

	}

	if len(res) > 1 {

		panic("Expression not valid")

	}

	res, x = res.Pop()
	return ConvertToInt(x)
}

func main() {
	flag.Parse()

	var raw = flag.Arg(0)

	tokenize(raw)

	var postfix = infixToPostfix(result)

	fmt.Println(eval(postfix))
}
