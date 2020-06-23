package chapter2

import "testing"

func TestFib3(t *testing.T) {
	// 0 1 2 3 4 5 6
	// 0 1 1 2 3 5 8
	if fib3(5) != 5 {
		t.Fatalf("failed, n: %d\n", 5)
	}

	if fib3(6) != 8 {
		t.Fatalf("failed, n: %d\n", 8)
	}
}
