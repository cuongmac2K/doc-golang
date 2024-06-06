package main

import "fmt"

func main() {
	b := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	c := removeNthFromEnd(&b, 2)
	c.show()
}
func (l *ListNode) show() {
	fmt.Println(l.Val)
	for l.Next != nil {
		l = l.Next
		fmt.Println(l.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := head, head
	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right != nil && right.Next != nil {
		left, right = left.Next, right.Next
	}

	left.Next = left.Next.Next
	return head
}
