package main

import "fmt"

func main() {
	list1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	sorted := mergeTwoLists(&list1, &list2)
	fmt.Println(sorted.Next.Next.Next.Next.Next.Next.Val)
}
