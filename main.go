package main

import "fmt"

func main() {
	nums1 := smallestEqual([]int{0,1,2})
	nums2 := smallestEqual([]int{4,3,2,1})
	nums3 := smallestEqual([]int{1,2,3,4,5,6,7,8,9,0})
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(nums3)
}
