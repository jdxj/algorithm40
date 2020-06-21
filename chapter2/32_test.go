package chapter2

import (
	"fmt"
	"testing"
)

func TestDFSQueens_SolveNQueens(t *testing.T) {
	dq := NewDFSQueens()

	result := dq.SolveNQueens(4)
	for _, v := range result {
		fmt.Printf("%v\n", v)
	}
}
