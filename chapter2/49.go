package chapter2

// 300 最长上升子序列
type LIS interface {
	LengthOfLIS(nums []int) int
}

// 1. 暴力, O(2^n)

// 2. DP O(n^2)
//   定义 DP[i]: 选择第 i 个元素时的最长上升子序列
//   Max(DP[0], ..., DP[i-1]) = DP[j]
//   DP[i] = DP[j] + 1
type DpLIS struct {
}

func (dl *DpLIS) LengthOfLIS(nums []int) (res int) {
	if len(nums) == 0 {
		return
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1 // 本身
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = maxC(dp[i], dp[j]+1)
			}
		}
		res = maxC(res, dp[i])
	}
	return
}

// 3. 贪心 + 二分搜索, O(nlogn)
