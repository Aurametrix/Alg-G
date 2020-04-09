type Stack []interface{}

func (s Stack) Peek() interface{} {
	return s[len(s)-1]
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) Push(value interface{}) {
	*s = append(*s, value)
}
