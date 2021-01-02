package _302_Deepest_Leaves_Sum


type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	maxLevel := 0

	var calculateSum func(node *TreeNode, level int)

	calculateSum = func(node *TreeNode, level int) {

		if level > maxLevel {
			sum = 0
			maxLevel = level
		}

		if level == maxLevel {
			sum += node.Val
		}

		if node.Left != nil {
			calculateSum(node.Left, level + 1)
		}

		if node.Right != nil {
			calculateSum(node.Right, level + 1)
		}

	}

	calculateSum(root, 0 )
	return sum

}
