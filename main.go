package main

import "fmt"

func main() {

	head := ListNode{
		Val :1,
		Next: &ListNode{
			Val :2,
			Next: &ListNode{
				Val :3,
				Next: &ListNode{
					Val:4,
					Next: &ListNode{
						Val :4,
						Next: &ListNode{
							Val:5,
							Next: &ListNode{
								Val:5,
								Next: &ListNode{
									Val:5,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	deleteDuplicates(&head)
	for head.Next != nil {
		fmt.Println(head.Val)
		head = *head.Next
	}
}
