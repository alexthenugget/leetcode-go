package problems

//Time complexity - O(n), memory - O(n)

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var stackP []*TreeNode
	var stackQ []*TreeNode
	findPath(root, p, &stackP)
	findPath(root, q, &stackQ)
	var result *TreeNode
	minLen := min(len(stackP), len(stackQ))
	for i := 0; i < minLen; i++ {
		if stackP[i] == stackQ[i] {
			result = stackP[i]
		} else {
			break
		}
	}
	return result
}

func findPath(node *TreeNode, target *TreeNode, stack *[]*TreeNode) bool {
	if node == nil {
		return false
	}
	(*stack) = append((*stack), node)
	if node == target {
		return true
	}
	if !findPath(node.Left, target, stack) && !findPath(node.Right, target, stack) {
		(*stack) = (*stack)[:len((*stack))-1]
		return false
	}
	return true
}
