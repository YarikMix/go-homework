package calc

import (
	"errors"
	"strconv"
)

func ConvertToFloat(x string) float64 {
	var res, _ = strconv.ParseFloat(x, 64)
	return res
}

func evalBinary(x string, y string, op string) float64 {
	if op == "+" {
		return ConvertToFloat(x) + ConvertToFloat(y)
	}

	if op == "-" {
		return ConvertToFloat(x) - ConvertToFloat(y)
	}

	if op == "*" {
		return ConvertToFloat(x) * ConvertToFloat(y)
	}

	if op == "/" {
		return ConvertToFloat(x) / ConvertToFloat(y)
	}

	return -1
}

func Eval(raw string) (float64, error) {

	var result = Tokenize(raw)

	var postfix, err = InfixToPostfix(result)

	if err != nil {
		return 0, errors.New("expression not valid")
	}

	var res Stack = make([]string, 0)

	var x, y string

	for _, token := range postfix {

		var _, ok = ops[token] // Если встретился оператор (+, -, *, /)

		if ok {

			res, y = res.Pop()
			res, x = res.Pop()

			res = res.Push(strconv.FormatFloat(evalBinary(x, y, token), 'E', -1, 64))

		} else { // Если встретилось число

			res = res.Push(token)

		}

	}

	if len(res) > 1 {
		return 0, errors.New("expression not valid")
	}

	res, x = res.Pop()
	return ConvertToFloat(x), nil
}
