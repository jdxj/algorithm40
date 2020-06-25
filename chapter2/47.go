package chapter2

// 152 乘积最大子序列
type Product interface {
	MaxProduct(nums []int) int
}

// 1. 暴力 Recursion 每个数字进行抉择: 乘, 不乘
type RecursionProduct struct {
	max int
}

func (rp *RecursionProduct) MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	rp.max = nums[0]
	rp.helper(nums, 1, nums[0], nums[0])
	return rp.max
}

func (rp *RecursionProduct) helper(nums []int, idx, max, min int) {
	if idx == len(nums) {
		return
	}

	curMax := maxC(max*nums[idx], maxC(min*nums[idx], nums[idx]))
	curMin := minC(max*nums[idx], minC(min*nums[idx], nums[idx]))
	rp.max = maxC(rp.max, curMax)
	rp.helper(nums, idx+1, curMax, curMin)
}

// 2. DP 两步曲
//   - 状态定义
//   - 状态转换

// (Max 在这里)
// DP[i][0] =  --
//            | DP[i-1][0] * a[i]     a[i] >= 0
//           -|
//            | DP[i-1][1] * a[i]     a[i] < 0
//             --
// DP[i][1] =  --
//            | DP[i-1][1] * a[i]     a[i] >= 0
//           -|
//            | DP[i-1][0] * a[i]     a[i] < 0
//             --

type DPProduct struct {
}

func (dpp *DPProduct) MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := [][]int{
		{nums[0], nums[0]},
		{0, 0},
	}
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		x, y := i%2, (i-1)%2
		// 规律: nums[i] 本身也需要参与比较, 其含义是: 前面的子序列的乘积没有 nums[i] 大.
		dp[x][0] = maxC(dp[y][0]*nums[i], maxC(dp[y][1]*nums[i], nums[i]))
		dp[x][1] = minC(dp[y][0]*nums[i], maxC(dp[y][1]*nums[i], nums[i]))
		res = maxC(res, dp[x][0])
	}
	return res
}

func maxC(p1, p2 int) int {
	if p1 > p2 {
		return p1
	}
	return p2
}

func minC(p1, p2 int) int {
	if p1 < p2 {
		return p1
	}
	return p2
}
