package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var sortedList ListNode

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &sortedList
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			dummy.Next = list1
			list1 = list1.Next
		} else {
			dummy.Next = list2
			list2 = list2.Next
		}

		dummy = dummy.Next
	}
	if list1 != nil {
		dummy.Next = list1
	} else {
		dummy.Next = list2
	}

	return sortedList.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	listsLen := len(lists)
	if listsLen == 0 {
		return nil
	}

	if listsLen == 1 {
		return lists[0]
	}

	right := mergeKLists(lists[:(listsLen+1)/2])
	left := mergeKLists(lists[(listsLen+1)/2:])

	sorted := mergeTwoLists(right, left)

	return sorted
}
