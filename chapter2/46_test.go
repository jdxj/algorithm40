package chapter2

import "testing"

func TestMinimumTotal(t *testing.T) {
	if minimumTotal(newTriangle()) != 11 {
		t.Fatalf("failed")
	}
}
