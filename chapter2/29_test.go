package chapter2

import "testing"

func TestDCCountDepth_MaxDepth(t *testing.T) {
	dc := &DCCountDepth{}

	if dc.MinDepth(BTree()) != 3 {
		t.Fatalf("failed")
	}
}
