package calc

type Stack []string

func (s Stack) Push(v string) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, string) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s Stack) Peek() string {
	l := len(s)
	return s[l-1]
}
