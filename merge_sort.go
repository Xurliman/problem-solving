package main

type Pair struct {
	Key   int
	Value string
}

//fmt.Println("mergeSort", mergeSort([]Pair{
//	{5, "apple"},
//	{2, "banana"},
//	{9, "cherry"},
//}))
//fmt.Println("mergeSort", mergeSort([]Pair{
//	{3, "cat"},
//	{3, "bird"},
//	{2, "dog"},
//}))

func mergeSort(pairs []Pair) []Pair {
	if len(pairs) <= 1 {
		return pairs
	}

	m := (len(pairs) + 1) / 2

	left := mergeSort(pairs[:m])
	right := mergeSort(pairs[m:])

	out := merge(left, right)

	return out
}

func merge(left, right []Pair) []Pair {
	l, r := 0, 0

	out := make([]Pair, 0, len(left)+len(right))
	for l < len(left) && r < len(right) {
		if left[l].Key > right[r].Key {
			out = append(out, right[r])
			r++
		} else {
			out = append(out, left[l])
			l++
		}
	}

	out = append(out, left[l:]...)
	out = append(out, right[r:]...)

	return out
}
