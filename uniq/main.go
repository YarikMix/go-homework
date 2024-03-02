package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"uniq/uniq"
)

func readLines(reader io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func writeLines(writer io.Writer, lines []string) error {
	w := bufio.NewWriter(writer)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}

	return w.Flush()
}

func Parse() uniq.Args {
	var args uniq.Args

	flag.BoolVar(&args.C, "c", false, "Подсчитать количество встречаний строки во входных данных. Вывести это число перед строкой отделив пробелом.")
	flag.BoolVar(&args.D, "d", false, "Вывести только те строки, которые повторились во входных данных.")
	flag.BoolVar(&args.U, "u", false, "Вывести только те строки, которые не повторились во входных данных.")
	flag.BoolVar(&args.I, "i", false, "Не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделённый пробелом.")
	flag.IntVar(&args.F, "f", 0, "Не учитывать первые num_chars символов в строке. При использовании вместе с параметром -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля).")
	flag.IntVar(&args.S, "s", 0, "Не учитывать регистр букв.")

	flag.Parse()

	return args
}

func main() {

	var args = Parse()

	if !args.IsValid() {
		flag.Usage()
		os.Exit(-1)
	}

	// Чтение данных

	var reader io.Reader
	reader = os.Stdin

	var input = flag.Arg(0)

	if input != "" {
		file, err := os.Open(input)

		if err != nil {
			log.Fatalf("readLines: %s", err)
		}

		reader = file

		defer file.Close()

	}

	lines, err := readLines(reader)

	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Обработка данных

	var arr = uniq.Solve(lines, args)

	// Вывод данных

	var writer io.Writer
	writer = os.Stdout

	var output = flag.Arg(2)

	if output != "" {
		file, err := os.Create(output)

		if err != nil {
			return
		}

		writer = file

		defer file.Close()

	}

	writeLines(writer, arr)
}
