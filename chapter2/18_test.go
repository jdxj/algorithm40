package chapter2

import (
	"testing"

	"github.com/jdxj/algorithm40/module"
)

//       3
//    /    \
//   1      5
//    \
//     2
func ValidBST1() *module.TreeNode {
	node3 := &module.TreeNode{Val: 3}
	node1 := &module.TreeNode{Val: 1}
	node5 := &module.TreeNode{Val: 5}
	node2 := &module.TreeNode{Val: 2}

	node3.Left = node1
	node3.Right = node5
	node1.Right = node2

	return node3
}

//    0
func ValidBST2() *module.TreeNode {
	node0 := &module.TreeNode{Val: 0}
	return node0
}

//       5
//    /    \
//   1      4
//        /  \
//       3    6
func InvalidBST() *module.TreeNode {
	node5 := &module.TreeNode{Val: 5}
	node1 := &module.TreeNode{Val: 1}
	node4 := &module.TreeNode{Val: 4}
	node3 := &module.TreeNode{Val: 3}
	node6 := &module.TreeNode{Val: 6}

	node5.Left = node1
	node5.Right = node4
	node4.Left = node3
	node4.Right = node6

	return node5
}

func TestInOrderVerifyBST_IsValidBST(t *testing.T) {
	iov := &InOrderVerifyBST{}

	if !iov.IsValidBST(ValidBST1()) {
		t.Fatalf("failed: 1")
	}

	if !iov.IsValidBST(ValidBST2()) {
		t.Fatalf("failed: 2")
	}

	if iov.IsValidBST(InvalidBST()) {
		t.Fatalf("failed: 3")
	}
}

func TestDefinitionVerifyBST_IsValidBST(t *testing.T) {
	dv := &DefinitionVerifyBST{}

	if !dv.IsValidBST(ValidBST1()) {
		t.Fatalf("failed: 1")
	}
}
