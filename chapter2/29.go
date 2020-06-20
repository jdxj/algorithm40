package chapter2

import "github.com/jdxj/algorithm40/module"

// 104/111 二叉树的最大/最小深度
type CountDepth interface {
	MinDepth(root *module.TreeNode) int
	MaxDepth(root *module.TreeNode) int
}

// 规律是找根到各叶子结点的最大/最小距离,

// 1. 递归
// 2. DFS
// 3. BFS

// 111
type DCCountDepth struct {
}

func (dc *DCCountDepth) MinDepth(root *module.TreeNode) int {
	if root == nil {
		return 0
	}

	leftMin := dc.MinDepth(root.Left)
	rightMin := dc.MinDepth(root.Right)

	if leftMin < rightMin {
		return leftMin + 1
	}
	return rightMin + 1
}

func (dc *DCCountDepth) MaxDepth(root *module.TreeNode) int {
	return 0
}
