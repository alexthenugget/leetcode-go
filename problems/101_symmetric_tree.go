package problems

//Time complexity - O(n), memory - O(n)

/*type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}*/

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var queue []*TreeNode
	queue = append(queue, root.Left)
	queue = append(queue, root.Right)
	for len(queue) > 0 {
		l := len(queue)
		j := l - 1
		for i := 0; i <= j; i++ {
			if queue[i] == nil && queue[j] != nil {
				return false
			}
			if queue[j] == nil && queue[i] != nil {
				return false
			}
			if queue[i] != nil && queue[j] != nil {
				if queue[i].Val != queue[j].Val {
					return false
				}
			}
			j--
		}
		for i := 0; i < l; i++ {
			if queue[i] != nil {
				queue = append(queue, queue[i].Left)
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[l:]
	}
	return true
}
