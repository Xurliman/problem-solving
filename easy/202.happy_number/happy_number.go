package _02_happy_number

// n = 19
func isHappy(n int) bool {
	var slow, fast = n, n

	for {
		slow = sumOfSquares(slow)
		fast = sumOfSquares(sumOfSquares(fast))

		if fast == 1 {
			return true
		}

		if slow == fast {
			return false
		}
	}
}

func sumOfSquares(n int) int {
	var sum int
	for n != 0 {
		sum += (n % 10) * (n % 10)

		n /= 10
	}

	return sum
}
