package chapter2

// 动态规划 (下)

// DP vs 回溯 vs 贪心
//   - 回溯 (递归): 重复计算
//   - 贪心: 永远局部最优
//   - DP: 记录局部最优子结构/多种记录值

func countPaths(grid [][]bool, row, col int) int {
	if !validSquare(grid, row, col) {
		return 0
	}
	if isAtEnd(grid, row, col) {
		return 1
	}
	return countPaths(grid, row+1, col) + countPaths(grid, row, col+1)
}

func validSquare(grid [][]bool, row, col int) bool {
	if grid[row][col] {
		return false
	}
	return true
}

func isAtEnd(grid [][]bool, row, col int) bool {
	if row == len(grid)-1 &&
		col == len(grid[0])-1 {
		return true
	}
	return false
}

// 优化:
//   - 记忆化
//   - 递推 (自下而上)

// opt[i][j] = opt[i+1][j] + opt[i][j+1]
func countPath2(grid [][]bool) (int, [][]int) {
	opt := newOpt(grid)
	rowNum := len(grid)
	colNum := len(grid[0])

	for row := rowNum - 2; row >= 0; row-- {
		for col := colNum - 2; col >= 0; col-- {
			if validSquare(grid, row, col) {
				opt[row][col] = opt[row+1][col] + opt[row][col+1]
			} else {
				opt[row][col] = 0
			}
		}
	}

	return opt[0][0], opt
}

func newGrid() [][]bool {
	grid := [][]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, true, false},
		{false, false, false, false, true, false, false, false},
		{true, false, true, false, false, true, false, false},
		{false, false, true, false, false, false, false, false},
		{false, false, false, true, true, false, true, false},
		{false, true, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false},
	}
	return grid
}

func newOpt(grid [][]bool) [][]int {
	rowNum := len(grid)
	colNum := len(grid[0])
	opt := make([][]int, rowNum)
	for i := range opt {
		opt[i] = make([]int, colNum)
	}
	for i := 0; i < rowNum; i++ {
		opt[i][colNum-1] = 1
	}
	for i := 0; i < colNum; i++ {
		opt[rowNum-1][i] = 1
	}
	return opt
}
