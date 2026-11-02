package problems

//Time complexity - O(n), memory - O(n)

func isBalanced(root *TreeNode) bool {
	if balancedHeight(root) == -1 {
		return false
	} else {
		return true
	}
}

func balancedHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := balancedHeight(node.Left)
	right := balancedHeight(node.Right)
	if left == -1 || right == -1 {
		return -1
	}
	if left-right > 1 || left-right < -1 {
		return -1
	}
	return max(left, right) + 1
}
