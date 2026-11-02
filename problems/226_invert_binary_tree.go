package problems

//Time complexity - O(n), memory - O(n)

func invertTree(root *TreeNode) *TreeNode {
	return invert(root)
}

func invert(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	left := invert(node.Left)
	right := invert(node.Right)
	node.Left = right
	node.Right = left
	return node
}
