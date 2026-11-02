package problems

//Time complexity - O(n), memory - O(1)

func isPalindromeList(head *ListNode) bool {
	curr := head
	curr2 := head
	for curr2.Next != nil && curr2.Next.Next != nil {
		curr = curr.Next
		curr2 = curr2.Next.Next
	}
	var currNode *ListNode
	var prevNode *ListNode
	if curr.Next != nil {
		currNode = curr.Next
		for currNode != nil {
			nextNode := currNode.Next
			currNode.Next = prevNode
			prevNode = currNode
			currNode = nextNode
		}
	}
	point := head
	curr = prevNode
	for curr != nil && point != nil {
		if point.Val != curr.Val {
			return false
		}
		point = point.Next
		curr = curr.Next
	}
	return true
}
