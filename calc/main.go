package main

import (
	"flag"
	"fmt"
	"unicode"
)

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

func main() {
	flag.Parse()

	var raw = flag.Arg(0)

	tokenize(raw)

	for _, str := range result {
		fmt.Println(str)
	}
}
