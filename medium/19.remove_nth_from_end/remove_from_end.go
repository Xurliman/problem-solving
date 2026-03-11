package _9_remove_nth_from_end

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func show() {
	res := removeNthFromEnd(&ListNode{
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

	for res != nil {
		fmt.Println("res", res)
		res = res.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var dummyPtr ListNode
	dummyPtr.Next = head

	var left, right *ListNode
	left = &dummyPtr
	right = head
	for n != 0 && right != nil {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return dummyPtr.Next
}
