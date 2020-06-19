package chapter2

import "github.com/jdxj/algorithm40/module"

// 235/236(二叉搜索树/二叉树) 的最近公共祖先
// 说明:
//    所有节点的值都是唯一的。
//    p、q 为不同节点且均存在于给定的二叉搜索树中。
type FindLowestCommonAncestor interface {
	// 感觉题目给 p, q 的指针不太合理, 应该使用 int
	LowestCommonAncestor(root, p, q *module.TreeNode) *module.TreeNode
}

// 寻找两个结点的最近公共祖先, 说明这两个结点分别在最近公共祖先的左右子树中

// 235 二叉搜索树
// 根据大小判断
type BSTreeFLCA struct {
}

func (bst *BSTreeFLCA) LowestCommonAncestor(root, p, q *module.TreeNode) *module.TreeNode {
	// 在 root 的左子树中
	if p.Val < root.Val && q.Val < root.Val {
		return bst.LowestCommonAncestor(root.Left, p, q)
	}
	// 在 root 的右子树中
	if p.Val > root.Val && q.Val > root.Val {
		return bst.LowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// 236 普通二叉树
type BTreeFLCA struct {
}

func (bt *BTreeFLCA) LowestCommonAncestor(root, p, q *module.TreeNode) *module.TreeNode {
	return bt.findPOrQ(root, p, q)
}

// findPorQ 在 root 中 (包括 root) 寻找 p 或者 q.
// 返回值不为 nil 代表 root 中包含 p 或者 q.
func (bt *BTreeFLCA) findPOrQ(root, p, q *module.TreeNode) *module.TreeNode {
	if root == nil ||
		root == p || root == q {
		return root // 就是个标记, 是否找到
	}

	// 在左子树中查找
	existLeft := bt.findPOrQ(root.Left, p, q)
	// 在右子树中查找
	existRight := bt.findPOrQ(root.Right, p, q)

	// 根据规律 (规律指: 当 p, q 分别在 root 两侧时, 那么 root 就是 "最近公共祖先"),
	// 找最近公共祖先就是判断 left, right 的情况 (找没找到 p, q).
	if existLeft == nil {
		return existRight // 根据说明: p,q 均存在, 所以可以直接返回
	} else {
		if existRight == nil {
			return existLeft
		}
		return root
	}
}
