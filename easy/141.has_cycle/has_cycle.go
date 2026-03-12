package main

type ListNode struct {
	Val  int
	Next *ListNode
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
