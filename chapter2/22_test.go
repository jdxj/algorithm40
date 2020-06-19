package chapter2

import (
	"math"
	"testing"
)

func TestIteratePow_MyPow(t *testing.T) {
	ip := &IteratePow{}

	if math.Pow(2, 8) != ip.MyPow(2, 8) {
		t.Fatalf("failed, x: %d, y: %d", 2, 8)
	}
	if math.Pow(2, 9) != ip.MyPow(2, 9) {
		t.Fatalf("failed, x: %d, y: %d", 2, 9)
	}
}
