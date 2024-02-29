package uniq

import (
	"fmt"
	"strings"
)

type Args struct {
	C bool
	D bool
	U bool
	I bool
	F int
	S int
}

func (args Args) IsValid() bool {
	if args.D && args.U || args.D && args.C || args.U && args.C {
		return false
	}

	return true
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

func cutWords(line string, num int) string {
	var array = strings.Fields(line)

	if len(array) < num {
		return ""
	}

	return strings.Join(array[num:], " ")
}

func cutRune(line string, num int) string {
	var array = []rune(line)

	if len(array) < num {
		return ""
	}

	return string(array[num:])
}

func getLine(line string, args Args) string {

	if args.I {
		line = strings.ToLower(line)
	}

	if args.F > 0 {
		line = cutWords(line, args.F)
	}

	if args.S > 0 {
		line = cutRune(line, args.S)
	}

	return line
}

func Compare(str1 string, str2 string, args Args) bool {
	return getLine(str1, args) == getLine(str2, args)
}

func Solve(lines []string, args Args) (arr []string) {

	var tmp string
	var c = 0
	var f = true

	for _, line := range lines {

		if args.D {

			if f && Compare(tmp, line, args) {
				arr = append(arr, line)
				f = false
			} else {
				f = true
			}

		} else if args.U {

			if count(line, lines) == 1 {
				arr = append(arr, line)
			}

		} else if !Compare(tmp, line, args) {

			if args.C {
				if c > 0 {
					s := fmt.Sprintf("%d %s", c, tmp)
					arr = append(arr, s)
				}
			} else {
				arr = append(arr, line)
			}

			c = 0
		}

		tmp = line
		c += 1
	}

	if args.C {
		s := fmt.Sprintf("%d %s", c, tmp)
		arr = append(arr, s)
	}

	return
}
