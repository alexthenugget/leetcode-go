package problems

//Time complexity - O(n), memory - O(1)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	prev := &ListNode{Next: head}
	first := prev
	second := prev
	for i := 0; i <= n; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return prev.Next
}
