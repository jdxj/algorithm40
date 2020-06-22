package chapter2

import (
	"fmt"
	"testing"
)

func TestNot(t *testing.T) {
	var a int = 10 // (1010)
	fmt.Printf("%d\n", a&(-a))

	a = 3 // (11)
	fmt.Printf("%d\n", a&(-a))

	a = 16 // (10000)
	fmt.Printf("%d\n", a&(-a))
}

func TestBitQueens_TotalNQueens(t *testing.T) {
	bq := &BitQueens{}
	result := bq.TotalNQueens(4)
	fmt.Printf("%d\n", result)

	result = bq.TotalNQueens(8)
	fmt.Printf("%d\n", result)
}
