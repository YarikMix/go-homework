package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
)

var flag_c = flag.Bool("c", false, "")
var flag_d = flag.Bool("d", false, "")
var flag_u = flag.Bool("u", false, "")
var flag_i = flag.Bool("i", false, "")
var flag_f = flag.Int("f", 0, "")
var flag_s = flag.Int("s", 0, "")

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

func count(str string, lines []string) int {
	var c = 0
	for _, line := range lines {
		if line == str {
			c += 1
		}
	}

	return c
}

func Split(str string) []string {
	return strings.Split(str, " ")
}

func Join(str []string) string {
	return strings.Join(str, " ")
}

func Compare(str1 string, str2 string) bool {
	return getLine(str1) == getLine(str2)
}

func getLine(line string) string {
	if *flag_s > 0 && len(line) > 0 && *flag_f > 0 && len(Split(line)) > 0 {
		if *flag_f >= len(Split(line)) {
			return ""
		}

		var res = Join(Split(line)[*flag_f:])[*flag_s:]

		if *flag_i {
			return strings.ToLower(res)
		}

		return res
	}

	if *flag_s > 0 && len(line) > 0 {

		if *flag_i {
			return strings.ToLower(line[*flag_s:])
		}

		return line[*flag_s:]
	}

	if *flag_f > 0 && len(Split(line)) > 0 {
		if *flag_f >= len(Split(line)) {
			return ""
		}

		var res = Join(Split(line)[*flag_f:])

		if *flag_i {
			return strings.ToLower(res)
		}

		return res
	}

	if *flag_i {
		return strings.ToLower(line)
	}

	return line
}

func test1(lines []string) (arr []string) {
	var tmp string = lines[0]

	arr = append(arr, lines[0])

	for _, line := range lines[1:] {

		if !Compare(tmp, line) {
			arr = append(arr, line)
		}

		tmp = line
	}

	return
}

func test2(lines []string) (arr []string) {
	var tmp string = lines[0]
	var c = 1

	for _, line := range lines[1:] {
		if !Compare(tmp, line) {
			s := fmt.Sprintf("%d %s", c, tmp)
			arr = append(arr, s)
			c = 0
		}

		tmp = line
		c += 1
	}

	s := fmt.Sprintf("%d %s", c, tmp)
	arr = append(arr, s)

	return
}

func test3(lines []string) (arr []string) {
	var prevLine string = lines[0]
	var f = true

	for _, line := range lines[1:] {
		if f && Compare(prevLine, line) && !slices.Contains(arr, line) {
			arr = append(arr, line)
			f = false
		} else {
			f = true
		}

		prevLine = line
	}

	return
}

func test4(lines []string) (arr []string) {
	for _, line := range lines {
		if count(line, lines) == 1 {
			arr = append(arr, line)
		}
	}

	return
}

func main() {

	flag.Parse()

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

	arr := make([]string, 0)
	if *flag_c {
		arr = test2(lines)
	} else if *flag_d {
		arr = test3(lines)
	} else if *flag_u {
		arr = test4(lines)
	} else {
		arr = test1(lines)
	}

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
