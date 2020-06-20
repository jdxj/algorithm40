package chapter2

import (
	"fmt"
	"testing"
)

func TestBFSLevelTraversal_LevelOrder(t *testing.T) {
	bfs := NewBFSLevelTraversal()

	result := bfs.LevelOrder(BTree())
	for level, data := range result {
		fmt.Printf("level: %d, data: %v\n", level, data)
	}
}
