package chapter2

import (
	"fmt"
	"testing"
)

func TestOpt(t *testing.T) {
	result := newOpt(newGrid())
	for _, v := range result {
		fmt.Printf("%v\n", v)
	}
}

func TestCountPath2(t *testing.T) {
	result, opt := countPath2(newGrid())
	if result != 27 {
		t.Fatalf("failed")
	}
	for _, v := range opt {
		fmt.Printf("%v\n", v)
	}
}
