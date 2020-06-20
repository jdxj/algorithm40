package chapter2

import (
	"fmt"
	"testing"
)

func TestRecursionBG_GenerateParenthesis(t *testing.T) {
	rbg := &RecursionBG{}

	result := rbg.GenerateParenthesis(10)
	for _, v := range result {
		fmt.Printf("%s\n", v)
	}
}
