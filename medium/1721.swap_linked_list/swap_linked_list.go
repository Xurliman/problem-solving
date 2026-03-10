package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	first := head
	for k > 1 {
		first = first.Next
		k--
	}

	prev, cur := head, first

	for cur.Next != nil {
		prev = prev.Next
		cur = cur.Next
	}

	prev.Val, first.Val = first.Val, prev.Val

	return head
}

func Show() {
	head := swapNodes(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 7,
								Next: &ListNode{
									Val:  8,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}, 4)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
