package main

import "fmt"

type ListNode struct {
	Val  string
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var prev *ListNode
	cur := slow

	for cur != nil {
		tmpNext := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmpNext
	}

	lead, lag := prev, head

	for lead != nil {
		if lead.Val != lag.Val {
			return false
		}

		lead = lead.Next
		lag = lag.Next
	}
	return true
}

func Show() {
	head := isPalindrome(&ListNode{
		Val: "A",
		Next: &ListNode{
			Val: "B",
			Next: &ListNode{
				Val: "C",
				Next: &ListNode{
					Val: "D",
					Next: &ListNode{
						Val: "D",
						Next: &ListNode{
							Val: "C",
							Next: &ListNode{
								Val: "B",
								Next: &ListNode{
									Val:  "D",
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	})
	fmt.Println(head)
}
