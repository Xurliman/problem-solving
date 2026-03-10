package _06_relative_ranks

import (
	"fmt"
	"sort"
)

const (
	first  = "Gold Medal"
	second = "Silver Medal"
	third  = "Bronze Medal"
)

func findRelativeRanks(scores []int) []string {
	result := make([]string, len(scores))
	spots := make(map[int]int)

	for k, v := range scores {
		spots[v] = k
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	for i := 0; i < len(scores); i++ {
		switch i {
		case 0:
			result[spots[scores[i]]] = first
		case 1:
			result[spots[scores[i]]] = second
		case 2:
			result[spots[scores[i]]] = third
		default:
			result[spots[scores[i]]] = fmt.Sprintf("%v", i+1)
		}
	}

	return result
}
