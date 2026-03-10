package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	f, s := head, head

	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	var prev *ListNode
	cur := s

	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}

	left := head
	right := prev

	for right != nil && right.Next != nil {
		tmpLeft := left.Next
		tmpRight := right.Next

		left.Next = right
		right.Next = tmpLeft

		left = tmpLeft
		right = tmpRight
	}
}

func Show() {
	reorderList(&ListNode{
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
	})

}
