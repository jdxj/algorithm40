package chapter2

import "testing"

func TestFib(t *testing.T) {
	// 0 1 2 3 4 5 6
	// 0 1 1 2 3 5 8
	result := fib(6, make(map[int]int))
	if result != 8 {
		t.Fatalf("failed\n")
	}
}

func TestFib2(t *testing.T) {
	result := fib2(6)
	if result != 8 {
		t.Fatalf("failed\n")
	}
}
