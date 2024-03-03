package calc

import (
	"errors"
	"unicode"
)

func emptyNumberBuffer(result []string, numberBuffer []rune) ([]string, []rune) {
	if len(numberBuffer) > 0 {
		result = append(result, string(numberBuffer))
		numberBuffer = nil
	}
	return result, numberBuffer
}

func emptyOperatorBuffer(result []string, operatorBuffer []rune) ([]string, []rune) {
	if len(operatorBuffer) > 0 {
		result = append(result, string(operatorBuffer))
		operatorBuffer = nil
	}
	return result, operatorBuffer
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

func Tokenize(raw string) ([]string, error) {
	var result = make([]string, 0)
	var numberBuffer = make([]rune, 0)
	var operatorBuffer = make([]rune, 0)

	var str = []rune(raw)

	for _, ch := range str {

		if isDigit(ch) {

			numberBuffer = append(numberBuffer, ch)

			if len(operatorBuffer) > 0 {
				result, operatorBuffer = emptyOperatorBuffer(result, operatorBuffer)
			}

		} else if isLeftParenthesis(ch) {

			if len(numberBuffer) > 0 {
				result, numberBuffer = emptyNumberBuffer(result, numberBuffer)
				result = append(result, "*")
			}

			if len(operatorBuffer) > 0 {
				result, operatorBuffer = emptyOperatorBuffer(result, operatorBuffer)
			}

			result = append(result, "(")

		} else if isRightParenthesis(ch) {

			result, numberBuffer = emptyNumberBuffer(result, numberBuffer)
			result = append(result, ")")

		} else {

			if isOperator(ch) {
				result, numberBuffer = emptyNumberBuffer(result, numberBuffer)
				operatorBuffer = append(operatorBuffer, ch)
			} else {
				return nil, errors.New("обнаружен неизвестный символ")
			}
		}

	}

	if len(numberBuffer) > 0 {
		result, numberBuffer = emptyNumberBuffer(result, numberBuffer)
	}

	if len(operatorBuffer) > 0 {
		result, operatorBuffer = emptyOperatorBuffer(result, operatorBuffer)
	}

	return result, nil
}
