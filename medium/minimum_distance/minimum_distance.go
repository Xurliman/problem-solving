package main

func minimumDistance(nums []int) int {
	var mindis = -1
	var tuples = make(map[int][]int)
	for i, num := range nums {
		tuples[num] = append(tuples[num], i)
	}

	for _, tuple := range tuples {
		if len(tuple) < 3 {
			continue
		}

		for i := 0; i < len(tuple)-2; i++ {
			res := 2 * (tuple[i+2] - tuple[i])
			if mindis == -1 {
				mindis = res
			}

			if res < mindis {
				mindis = res
			}
		}
	}
	return mindis
}

func calc(i, j, k int) int {
	return 2 * (max(i, j, k) - min(i, j, k))
}
