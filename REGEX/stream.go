func MakeWordTotals() *WordTotals {
	return &WordTotals{ make(map[string]uint64) }
}

type WordTotals struct {
	WordTotals map[string]uint64
}

func (wordTotals *WordTotals) Update(word string) {
	total, found := wordTotals.WordTotals[word]
	if !found {
		total = 0
	}
	wordTotals.WordTotals[word] = total + 1
}

func (wordTotals *WordTotals) GetCount(word string) *WordCount {
	return &WordCount{word, wordTotals.WordTotals[word]}
}
