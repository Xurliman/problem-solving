package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := &ListNode{Val: 1}
	fmt.Println(hasCycle(a))
}

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

func hasCycle(head *ListNode) bool {
	var hash = make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}

		hash[head] = struct{}{}
		head = head.Next
	}

	return false
}
