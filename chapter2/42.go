package chapter2

// 52 N皇后
type QueensII interface {
	TotalNQueens(n int) int
}

// 1. 位运算 (最快)
// 用位运算实现
type BitQueens struct {
	limit int
	count int
}

func (bq *BitQueens) TotalNQueens(n int) int {
	bq.limit = n
	bq.count = 0
	bq.helper(0, 0, 0, 0)
	return bq.count
}

func (bq *BitQueens) helper(row, col, pie, na int) {
	// 终止判断
	if row >= bq.limit {
		bq.count++
		return
	}

	// 得到空位
	// (^(col | pie | na))
	//     0: 没被占; 1: 被占 -> 取反: 将 "没被占" 变为1
	// ((1 << bq.limit) - 1)
	//     四皇后 -> int32: 0...000001111
	//     八皇后 -> int32: 0...011111111; 因为前一部分使用了 "^", 所以变成了1, 按需要应该使其变为0; (只关心后 n 位)
	bit := (^(col | pie | na)) & ((1 << bq.limit) - 1)
	for bit > 0 {
		// 得到一个空位 ("最低位是1" 到 "最低位" 所组成的数: 1010(2) -> 10(2))
		p := bit & (-bit)
		// 占位
		bq.helper(row+1, col|p, (pie|p)<<1, (na|p)>>1)
		// 打掉最后一位1
		bit = bit & (bit - 1)
	}
}
