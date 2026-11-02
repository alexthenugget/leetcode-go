package problems

//Time complexity - O(n), memory - O(1)

func reverseList(head *ListNode) *ListNode {
	currNode := head
	var prevNode *ListNode
	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}
	return prevNode
}
