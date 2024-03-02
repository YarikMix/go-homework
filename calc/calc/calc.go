package calc

import "strconv"

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

func Eval(raw string) int {

	var result = Tokenize(raw)

	var postfix = InfixToPostfix(result)

	var res Stack = make([]string, 0)

	var x, y string

	for _, token := range postfix {

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
