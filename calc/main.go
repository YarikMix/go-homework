package main

import (
	"flag"
	"fmt"
	. "main/calc"
)

func main() {
	flag.Parse()

	var raw = flag.Arg(0)

	var result = Eval(raw)

	fmt.Println(result)
}
