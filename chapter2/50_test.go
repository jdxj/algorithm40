package chapter2

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	coins := []int{1, 4, 9}
	res := coinChange(coins, 12)
	fmt.Printf("%d\n", res)
}
