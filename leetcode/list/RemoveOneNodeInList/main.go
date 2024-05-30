package main

func main() {
	head1 := ListNode{1}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
head ListNode
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var count int
	curr := head
	for curr != nil {
		curr = curr.Next
		count++
	}

	curr = head
	for i := 1; i < count-n; i++ {
		curr = curr.Next
	}

	if count == n {
		head = head.Next
	} else if curr.Next != nil {
		curr.Next = curr.Next.Next
	}

	return head
}
