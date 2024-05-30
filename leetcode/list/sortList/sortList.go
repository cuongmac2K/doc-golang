package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil // key: unlinks the nodes in the middle

	left := sortList(head)
	right := sortList(slow)

	return mergeList(left, right)
}

func mergeList(l1, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	nn := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			nn.Next = l1
			l1 = l1.Next
		} else {
			nn.Next = l2
			l2 = l2.Next
		}
		nn = nn.Next
	}

	if l1 != nil {
		nn.Next = l1
	} else if l2 != nil {
		nn.Next = l2
	}

	return dummy.Next
}
