package chapter2

// 51/52 N 皇后

// 1. DFS
//   - 暴力遍历
//   - 记录不能放的标记

// 51
type Queens interface {
	SolveNQueens(n int) [][]string
}

func NewDFSQueens() *DFSQueens {
	dq := &DFSQueens{
		col: make(map[int]struct{}),
		pie: make(map[int]struct{}),
		naa: make(map[int]struct{}),
	}
	return dq
}

type DFSQueens struct {
	limit  int
	result [][]int

	col map[int]struct{}
	pie map[int]struct{}
	naa map[int]struct{}
}

func (dq *DFSQueens) SolveNQueens(n int) [][]string {
	dq.limit = n
	dq.helper(0, nil)
	return dq.genLayout()
}

// row: 在 row 行进行搜索
// cols: 每一行中 Q 放置的列号
// 规律:
//   - 列, 撇, 捺 标记技巧
//   - 清除标记
func (dq *DFSQueens) helper(row int, cols []int) {
	// 一种布局结束
	if row == dq.limit {
		// 所有行都放完, 收集结果
		dq.result = append(dq.result, cols)
		return
	}

	// 查看当前行 (row) 的哪一列 (col) 能放
	for col := 0; col < dq.limit; col++ {
		_, existCol := dq.col[col]     // 检查列是否能放
		_, existPie := dq.pie[row+col] // 检查撇是否能放
		_, existNaa := dq.naa[row-col] // 检查捺是否能放

		// 不能放
		if existCol || existPie || existNaa {
			continue
		}

		// 放置
		dq.col[col] = struct{}{}
		dq.pie[row+col] = struct{}{}
		dq.naa[row-col] = struct{}{}

		// 进入下一行
		dq.helper(row+1, append(cols, col))

		// 恢复 (如果当前行有多个空位, 需要清除之前的状态标记)
		delete(dq.col, col)
		delete(dq.pie, row+col)
		delete(dq.naa, row-col)
	}
}

func (dq *DFSQueens) genLayout() [][]string {
	var result [][]string
	// 先生成 "."
	var a string
	for i := 0; i < dq.limit; i++ {
		a += "."
	}
	for _, layouts := range dq.result {
		var tmp []string
		for _, col := range layouts {
			// 再用 "Q" 替换 "."
			b := a[:col] + "Q" + a[col+1:] + "\n"
			tmp = append(tmp, b)
		}
		result = append(result, tmp)
	}
	return result
}
