package chapter2

import "testing"

func TestDpLIS_LengthOfLIS(t *testing.T) {
	dl := &DpLIS{}
	correct := 4
	if num := dl.LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}); num != correct {
		t.Fatalf("correct: %d, num: %d\n", correct, num)
	}
}
