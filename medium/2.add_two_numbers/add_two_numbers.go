package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//
//var inMind int
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	if l1 == nil && l2 == nil {
//		if inMind != 0 {
//			return &ListNode{inMind, nil}
//		}
//		return nil
//	}
//
//	var cur ListNode
//	if l1 == nil {
//		sum := inMind + l2.Val
//		cur.Val = sum % 10
//		inMind = sum / 10
//		cur.Next = addTwoNumbers(nil, l2.Next)
//	} else if l2 == nil {
//		sum := inMind + l1.Val
//		cur.Val = sum % 10
//		inMind = sum / 10
//		cur.Next = addTwoNumbers(l1.Next, nil)
//	} else {
//		sum := inMind + l1.Val + l2.Val
//		cur.Val = sum % 10
//		inMind = sum / 10
//		cur.Next = addTwoNumbers(l1.Next, l2.Next)
//	}
//	return &cur
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addRecursive(l1, l2, 0)
}

func addRecursive(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}

	sum := carry
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	node := &ListNode{
		Val:  sum % 10,
		Next: addRecursive(l1, l2, sum/10),
	}
	return node
}
