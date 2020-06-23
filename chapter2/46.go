package chapter2

// 120 三角形最小路径和

// 1. 递归/回溯
// 2. 贪心 (不合适)

// 3. 动态规划
//    从下往上想
// DP[i, j]: bottom 到 [i, j] 的最小值
//   dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
//   dp 初始状态: dp[m-1][j] = triangle[m-1][j]
func minimumTotal(triangle [][]int) int {
	rowNum := len(triangle)
	min := make([]int, len(triangle[rowNum-1]))
	copy(min, triangle[rowNum-1])

	for row := rowNum - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			min[col] = triangle[row][col] + less(min[col], min[col+1])
		}
	}
	return min[0]
}

func less(p1, p2 int) int {
	if p1 < p2 {
		return p1
	}
	return p2
}

func newTriangle() [][]int {
	result := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	return result
}
