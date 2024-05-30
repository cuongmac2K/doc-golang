package main

import "fmt"

func main() {
	a := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: nil,
				},
			},
		},
	}
	fmt.Println(hasCycle(a))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head.Next
	for fast.Next != nil && fast.Next.Next != nil && fast != slow {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow == fast
}
