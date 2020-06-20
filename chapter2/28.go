package chapter2

import (
	"container/list"

	"github.com/jdxj/algorithm40/module"
)

// 102 层遍历二叉树
type LevelTraversal interface {
	LevelOrder(root *module.TreeNode) [][]int
}

func NewBFSLevelTraversal() *BFSLevelTraversal {
	bfs := &BFSLevelTraversal{
		list: list.New(),
	}
	return bfs
}

// 1. BFS O(n)
type BFSLevelTraversal struct {
	list *list.List
}

func (bfs *BFSLevelTraversal) LevelOrder(root *module.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int

	lis := bfs.list
	// 初始化第一个
	lis.PushBack(root)

	// 队列中仍有值, 则继续
	for lis.Front() != nil {
		var levelData []int   // 记录当前层的数据
		levelLen := lis.Len() // 当前层中数据的个数

		for i := 0; i < levelLen; i++ { // 取出对应个数元素
			// 获取第一个结点
			node := lis.Remove(lis.Front()).(*module.TreeNode)
			// 记录数据
			levelData = append(levelData, node.Val)
			// 入队子结点
			if node.Left != nil {
				lis.PushBack(node.Left)
			}
			if node.Right != nil {
				lis.PushBack(node.Right)
			}
		}
		result = append(result, levelData)
	}

	return result
}

// 2. DFS O(n)
//   - 传递 level
