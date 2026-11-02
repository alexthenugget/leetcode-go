package problems

import "math"

//Time complexity - O(n), memory - O(n)

func isValidBST(root *TreeNode) bool {
	return isValidNode(root, -int(math.Pow(2, 32)), int(math.Pow(2, 32)+1))
}

func isValidNode(node *TreeNode, leftBourder int, rightBourder int) bool {
	if node == nil {
		return true
	}
	if node.Val <= leftBourder || node.Val >= rightBourder {
		return false
	}
	return isValidNode(node.Left, leftBourder, node.Val) && isValidNode(node.Right, node.Val, rightBourder)
}
