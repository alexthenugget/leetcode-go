package problems

//Time complexity - O(n), memory - O(n)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	result = append(result, root.Val)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			right := node.Right
			left := node.Left
			if right != nil {
				queue = append(queue, right)
			}
			if left != nil {
				queue = append(queue, left)
			}
		}
		if len(queue) > 0 {
			result = append(result, queue[0].Val)
		}
	}
	return result
}
