package neetcode

type Pair struct {
	Key   int
	Value string
}

func insertionSort(pairs []Pair) [][]Pair {
	if len(pairs) == 0 {
		return [][]Pair{}
	}
	output := make([][]Pair, 0, len(pairs))

	for i := 1; i < len(pairs); i++ {
		j := i - 1
		p := make([]Pair, len(pairs))
		copy(p, pairs)
		output = append(output, p)
		for j >= 0 && pairs[j+1].Key < pairs[j].Key {
			pairs[j+1].Key, pairs[j].Key = pairs[j].Key, pairs[j+1].Key
			pairs[j+1].Value, pairs[j].Value = pairs[j].Value, pairs[j+1].Value
			j--
		}
	}
	p := make([]Pair, len(pairs))
	copy(p, pairs)
	output = append(output, p)
	return output
}
