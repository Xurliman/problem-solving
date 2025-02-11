package main

type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		if current.Next == nil {
			return current
		}

		if current.Val == current.Next.Val {
			current = current.Next
		} else {
			current = current.Next
		}
	}
	return head
}
