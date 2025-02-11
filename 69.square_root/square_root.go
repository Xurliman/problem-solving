package main

func mySqrt(x int) int {
	y := x >> 1
	var prev int
	for i:= 2; i <= y; i++ {
		if prev * prev > x {
			return i
		}
		prev = i
	}
	return prev
}