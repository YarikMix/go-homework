package calc

import (
	"errors"
	"strconv"
)

func ConvertToFloat(x string) (float64, error) {
	return strconv.ParseFloat(x, 64)
}

func evalBinary(x string, y string, op string) (float64, error) {
	var num1, err1 = strconv.ParseFloat(x, 64)
	if err1 != nil {
		return 0, errors.New("не удалось спарсить выражение")
	}

	var num2, err2 = strconv.ParseFloat(y, 64)
	if err2 != nil {
		return 0, errors.New("не удалось спарсить выражение")
	}

	if op == "+" {
		return num1 + num2, nil
	}

	if op == "-" {
		return num1 - num2, nil
	}

	if op == "*" {
		return num1 * num2, nil
	}

	if op == "/" {
		return num1 / num2, nil
	}

	return 0, errors.New("не удалось спарсить выражение")
}

func Eval(raw string) (float64, error) {

	var result, tokenizeError = Tokenize(raw)

	if tokenizeError != nil {
		return 0, tokenizeError
	}

	var postfix, err = InfixToPostfix(result)

	if err != nil {
		return 0, err
	}

	var res Stack = make([]string, 0)

	var x, y string

	for _, token := range postfix {

		var _, ok = ops[token] // Если встретился оператор (+, -, *, /)

		if ok {

			if len(res) < 2 {
				return 0, errors.New("не получилось спарсить выражение")
			}

			res, y = res.Pop()
			res, x = res.Pop()

			var tmp, err = evalBinary(x, y, token)

			if err != nil {
				return 0, err
			}

			res = res.Push(strconv.FormatFloat(tmp, 'E', -1, 64))

		} else { // Если встретилось число

			res = res.Push(token)

		}

	}

	if len(res) > 1 {
		return 0, errors.New("не получилось спарсить выражение")
	}

	res, x = res.Pop()

	return ConvertToFloat(x)
}
