package chapter2

import "github.com/jdxj/algorithm40/module"

// 98 验证二叉搜索树
type VerifyBST interface {
	IsValidBST(root *module.TreeNode) bool
}

// 1. 递归
// 中序遍历的结果是升序的. 这里只保存前驱节点的值, 而不必保存中序遍历的完整结果.
// O(n)
type InOrderVerifyBST struct {
	prev *module.TreeNode
}

func (iov *InOrderVerifyBST) IsValidBST(root *module.TreeNode) bool {
	defer func() {
		iov.prev = nil
	}()
	return iov.helper(root)
}

func (iov *InOrderVerifyBST) helper(root *module.TreeNode) bool {
	// 空树也是 bst
	if root == nil {
		return true
	}

	// 检查左子树是不是升序的
	if !iov.helper(root.Left) {
		return false
	}
	// 检查 '当前节点' 与 '中序遍历的前驱节点' 是否是升序的
	if iov.prev != nil && iov.prev.Val >= root.Val {
		return false
	}
	// 移动 prev
	iov.prev = root
	// 检查右子树是不是升序的
	return iov.helper(root.Right)
}

// 2. 递归
// 利用定义: 左子树(max) <= 当前节点 < 右子树(min)
// O(n)
type DefinitionVerifyBST struct {
}

func (dv *DefinitionVerifyBST) IsValidBST(root *module.TreeNode) bool {
	return dv.helper(root, nil, nil)
}

func (dv *DefinitionVerifyBST) helper(root *module.TreeNode, lower, upper *int) bool {
	// 空树也是 bst
	if root == nil {
		return true
	}

	// 不满足的情况返回 false
	if lower != nil && root.Val <= *lower {
		return false
	}
	if upper != nil && root.Val >= *upper {
		return false
	}

	// 往左走: 更新上界
	// 往右走: 更新下界
	return dv.helper(root.Left, lower, &root.Val) &&
		dv.helper(root.Right, &root.Val, upper)
}
