package main

import "fmt"

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}

func findLHS(nums []int) int {
	var res int
	for L := 0; L < len(nums); L++ {
		var count int
		for R := L + 1; R < len(nums); R++ {
			if nums[L]-nums[R] <= 1 && nums[L]-nums[R] >= -1 {
				count++
			} else {
				continue
			}

			if count > res {
				res = count
			}
		}
	}
	return res
}
