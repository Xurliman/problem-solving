package climbing_stairs

var l = map[int]int{
	1:1,
	2:2,
}

func climbStairs(n int) int {
	if i, ok := l[n]; ok {
		return i
	}
	left, right := climbStairs(n-1), climbStairs(n-2)
	l[n-1] = left
	l[n-2] = right
	l[n] = left + right
	return left + right
}
