package main

func bitwiseComplement(n int) int {
	zero := n
	var bitLength int
	for zero != 0 {
		zero = zero >> 1
		bitLength++
	}

	mask := (1 << bitLength) - 1

	return n ^ mask
}
