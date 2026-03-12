package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var ptrA, ptrB *ListNode

	ptrA = headA
	ptrB = headB

	var lenACovered, lenBCovered bool

	for {
		if ptrA == ptrB {
			return ptrA
		}

		if ptrA != nil {
			ptrA = ptrA.Next
		} else if !lenBCovered {
			ptrA = headB
			lenBCovered = true
		}

		if ptrB != nil {
			ptrB = ptrB.Next
		} else if !lenACovered {
			ptrB = headA
			lenACovered = true
		}
	}

}
