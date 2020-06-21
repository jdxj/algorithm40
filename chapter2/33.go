package chapter2

// 36/37 判断数独有效
// 搜索 + 剪枝

// 1. DFS
//   row[9]
//   col[9]
//   block[3][3] ([i/3][j/3])
//   check: row, col block
// 剪枝:
//   - 先填空格少的 row/col
//   - 预处理
