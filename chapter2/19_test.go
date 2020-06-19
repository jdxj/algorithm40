package chapter2

import (
	"testing"

	"github.com/jdxj/algorithm40/module"
)

//         6
//     /      \
//    2         8
//  /  \      /  \
// 0    4    7    9
//    /  \
//   3    5
func BTree() *module.TreeNode {
	node6 := &module.TreeNode{Val: 6}
	node2 := &module.TreeNode{Val: 2}
	node8 := &module.TreeNode{Val: 8}
	node0 := &module.TreeNode{Val: 0}
	node4 := &module.TreeNode{Val: 4}
	node7 := &module.TreeNode{Val: 7}
	node9 := &module.TreeNode{Val: 9}
	node3 := &module.TreeNode{Val: 3}
	node5 := &module.TreeNode{Val: 5}

	node6.Left = node2
	node6.Right = node8

	node2.Left = node0
	node2.Right = node4

	node8.Left = node7
	node8.Right = node9

	node4.Left = node3
	node4.Right = node5

	return node6
}

func TestBTreeFLCA_LowestCommonAncestor(t *testing.T) {
	bt := &BTreeFLCA{}

	tree := BTree()
	if bt.LowestCommonAncestor(tree, tree.Left, tree.Right) != tree {
		t.Fatalf("failed")
	}
}

func TestBSTreeFLCA_LowestCommonAncestor(t *testing.T) {
	bst := &BSTreeFLCA{}

	tree := BTree()
	if bst.LowestCommonAncestor(tree, tree.Right.Left, tree.Right.Right) != tree.Right {
		t.Fatalf("failed")
	}
}
