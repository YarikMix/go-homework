package main

import (
	"fmt"
	"main/calc"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Не передано выражение")
		os.Exit(-1)
	}

	var raw = os.Args[1]

	var result, _ = calc.Eval(raw)

	fmt.Println(result)
}
