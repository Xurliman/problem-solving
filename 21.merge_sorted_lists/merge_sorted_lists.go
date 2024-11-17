package merge_sorted_lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

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

	if list1.Val < list2.Val {
		sortedList.Val = list1.Val
		sortedList.Next = mergeTwoLists(list1.Next, list2)
	} else {
		sortedList.Val = list2.Val
		sortedList.Next = mergeTwoLists(list1, list2.Next)
	}
	return &sortedList
}
