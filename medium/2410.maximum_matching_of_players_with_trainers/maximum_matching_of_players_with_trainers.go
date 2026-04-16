package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)

	var matches int
	var l, r int

	for l < len(players) && r < len(trainers) {
		if players[l] <= trainers[r] {
			matches++
			l++
			r++
		} else {
			r++
		}
	}
	return matches
}
