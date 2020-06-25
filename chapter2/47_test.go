package chapter2

import "testing"

func TestDPProduct_MaxProduct(t *testing.T) {
	dpp := &DPProduct{}

	if num := dpp.MaxProduct([]int{2, 3, -2, 4}); num != 6 {
		t.Fatalf("failed: %d, num: %d\n", 6, num)
	}
	if dpp.MaxProduct([]int{-2, 0, -1}) != 0 {
		t.Fatalf("failed: %d\n", 0)
	}
	if dpp.MaxProduct([]int{0, 2}) != 2 {
		t.Fatalf("failed: %d\n", 2)
	}
	if dpp.MaxProduct([]int{-4, -3, -2}) != 12 {
		t.Fatalf("failed: %d\n", 12)
	}
}

func TestRecursionProduct_MaxProduct(t *testing.T) {
	rp := &RecursionProduct{}
	if num := rp.MaxProduct([]int{2, 3, -2, 4}); num != 6 {
		t.Fatalf("failed: %d, num: %d\n", 6, num)
	}
	if num := rp.MaxProduct([]int{-2, 0, -1}); num != 0 {
		t.Fatalf("failed: %d, num: %d\n", 0, num)
	}
	if num := rp.MaxProduct([]int{-4, -3, -2}); num != 12 {
		t.Fatalf("failed: %d, num: %d\n", 12, num)
	}
}
