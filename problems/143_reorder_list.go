package problems

//Time complexity - O(n), memory - O(1)

func reorderList(head *ListNode) {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode
	next := slow.Next
	slow.Next = nil
	slow = next
	for slow != nil {
		next = slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	first := head
	last := prev

	for first != nil && last != nil {
		nextF := first.Next
		nextL := last.Next

		first.Next = last
		last.Next = nextF

		first = nextF
		last = nextL
	}
}
