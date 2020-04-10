type Stack(type T) []T

func (s Stack(T)) Peek() T {
	return s[len(s)-1]
}

func (s *Stack(T)) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack(T)) Push(value T) {
	*s = append(*s, value)
}
