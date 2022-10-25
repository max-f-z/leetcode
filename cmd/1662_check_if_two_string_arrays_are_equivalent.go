package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	op1 := &operator{
		currentIdx:     0,
		currentWordIdx: 0,
		words:          word1,
	}
	op2 := &operator{
		currentIdx:     0,
		currentWordIdx: 0,
		words:          word2,
	}

	for {
		if op1.Current() != op2.Current() {
			return false
		}

		hasNext1 := op1.Next()
		hasNext2 := op2.Next()

		if hasNext1 != hasNext2 {
			return false
		}

		if !hasNext1 && !hasNext2 {
			break
		}
	}

	return true
}

type operator struct {
	currentIdx     int
	currentWordIdx int
	words          []string
}

func (op *operator) Next() bool {
	op.currentIdx++
	if op.currentIdx >= len(op.words[op.currentWordIdx]) {
		op.currentWordIdx++
		op.currentIdx = 0
		if op.currentWordIdx >= len(op.words) {
			return false
		}
	}

	return true
}

func (op *operator) Current() byte {
	return op.words[op.currentWordIdx][op.currentIdx]
}
