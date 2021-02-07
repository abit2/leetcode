package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	robbed, notRobbed := map[*TreeNode]int{}, map[*TreeNode]int{}
	return helper(root, false, robbed, notRobbed)
}

func helper(root *TreeNode, parentRobbed bool, robbed, notRobbed map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	var ok bool

	if parentRobbed == true {
		_, ok = notRobbed[root]
		if !ok {
			notRobbed[root] = helper(root.Left, false, robbed, notRobbed) + helper(root.Right, false, robbed, notRobbed)
		}

		return notRobbed[root]
	} else {
		_, ok = robbed[root]
		if !ok {
			robbed[root] = root.Val + helper(root.Left, true, robbed, notRobbed) + helper(root.Right, true, robbed, notRobbed)
		}

		_, ok = notRobbed[root]
		if !ok {
			notRobbed[root] = helper(root.Left, false, robbed, notRobbed) + helper(root.Right, false, robbed, notRobbed)
		}

		return max(robbed[root], notRobbed[root])
	}
	return 0
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
