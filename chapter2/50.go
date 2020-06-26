package chapter2

// 322 零钱兑换

// 1. 暴力
// 2. 贪心 (不合适)
// 3. DP
//   定义 DP[i]: 上到第 i 阶, 最少需要多少步 (到达金额 i 所需最少硬币数)
//   DP[i] = Min{DP[i - coins[j]} + 1
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 初始化一个最大硬币数
	for i := range dp {
		dp[i] = amount + 1 // +1 用于避免实际刚好需要 amount 个硬币
	}
	// 达到0金额所需最少硬币数为0
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		// 尝试每种硬币, 如果想要到达金额 i, 且使用第 j 种硬币的话,
		// 所需最少硬币数为 dp[i-coins[j]] + 1 (递推关系)
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i { // i 代表要满足的钱, 减去 coins[j] 时要大于等于0
				dp[i] = minC(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
