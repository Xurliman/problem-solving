package not_finished

import "fmt"

func Show() {
	mylist := Constructor()
	commands := []string{"addAtHead", "addAtHead", "addAtHead", "addAtIndex", "deleteAtIndex", "addAtHead", "addAtTail", "get", "addAtHead", "addAtIndex", "addAtHead"}
	values := [][]int{{7}, {2}, {1}, {3, 0}, {2}, {6}, {4}, {4}, {4}, {5, 0}, {6}}

	for idx, c := range commands {
		switch c {
		case "addAtHead":
			mylist.AddAtHead(values[idx][0])
		case "addAtTail":
			mylist.AddAtTail(values[idx][0])
		case "addAtIndex":
			mylist.AddAtIndex(values[idx][0], values[idx][1])
		case "get":
			fmt.Println(mylist.Get(values[idx][0]))
		case "deleteAtIndex":
			mylist.DeleteAtIndex(values[idx][0])
		}
	}

}

type MyNode struct {
	Val  int
	Next *MyNode
	Prev *MyNode
}

type MyLinkedList struct {
	Left  *MyNode
	Right *MyNode
}

func Constructor() MyLinkedList {
	left := &MyNode{
		Val: 0,
	}
	right := &MyNode{
		Val: 0,
	}

	left.Next = right
	right.Prev = left
	return MyLinkedList{
		Left:  left,
		Right: right,
	}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.Left.Next
	for cur != nil && index > 0 {
		index--
		cur = cur.Next
	}

	if cur != nil && index == 0 && cur != this.Right.Next {
		return cur.Val
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	node, next, prev := &MyNode{Val: val}, this.Left.Next, this.Left
	prev.Next = node
	next.Prev = node
	node.Prev = prev
	node.Next = next
}

func (this *MyLinkedList) AddAtTail(val int) {
	node, next, prev := &MyNode{Val: val}, this.Right, this.Right.Prev
	prev.Next = node
	next.Prev = node
	node.Prev = prev
	node.Next = next
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this.Left.Next
	for cur != nil && index > 0 {
		index--
		cur = cur.Next
	}

	if cur != nil && index == 0 {
		node, next, prev := &MyNode{Val: val}, cur, cur.Prev
		prev.Next = node
		next.Prev = node
		node.Prev = prev
		node.Next = next
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.Left.Next
	for cur != nil && index > 0 {
		index--
		cur = cur.Next
	}

	if cur != nil && index == 0 && cur != this.Right {
		prev := cur.Prev
		next := cur.Next
		prev.Next = next
		next.Prev = prev
	}
}
