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

	var result, err = calc.Eval(raw)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	fmt.Println(result)
}
