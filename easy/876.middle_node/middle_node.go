package easy

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var locs []*ListNode

	for head != nil {
		locs = append(locs, head)

		head = head.Next
	}

	return locs[len(locs)/2]
}

func memoryEfficientMiddleNode(head *ListNode) *ListNode {
	s, f := head, head

	for f != nil && f.Next != nil {
		s = head.Next
		f = head.Next.Next
	}

	return s
}

func Show() {
	fmt.Println(memoryEfficientMiddleNode(
		&ListNode{
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
		}),
	)
}
