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
	c := swapPairs(&b)
	c.show()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	head, head.Next, head.Next.Next = head.Next, swapPairs(head.Next.Next), head

	return head
}
func (l *ListNode) show() {
	fmt.Println(l.Val)
	for l.Next != nil {
		l = l.Next
		fmt.Println(l.Val)
	}
}
func swapPairs1(head *ListNode) *ListNode {
	preHead := &ListNode{Next: head}
	curr := preHead
	for curr.Next != nil && curr.Next.Next != nil {
		nextNext := curr.Next.Next
		next := curr.Next

		next.Next = nextNext.Next
		nextNext.Next = next

		curr.Next = nextNext
		curr = next
	}

	return preHead.Next
}
