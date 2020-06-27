package chapter2

// 72 编辑距离
// 操作: 插入, 删除, 替换

// 1. 暴力 (BFS)
// 2. DP
//   定义 DP[i][j]: word1 的前 i 个字符转换到 word2 的前 j 个字符最少步数.
//   转换:
//                 --
//                | DP[i-1][j-1]                                  word1[i] == word2[j]
//                |                                               ( 两个字符相等不用操作,
//                |                                                 也就是说 word1 的前 i 个字符转换成 word2 的前 j 个字符所需的
//                |                                                 最少步数就是 word1 的前 i-1 个字符转换成 word2 的前 j-1 个
//                |                                                 字符
//                |                                                 所需最少步数 )
//     D[i][j] = -|
//                | Min(DP[i-1][j], DP[i][j-1], DP[i-1][j-1]) + 1 word1[i] != word2[j]
//                |     插入             删除        替换
//                 --
// 规律: 该题可以理解为将任意字符串 word1 转换乘 word2 所需最少操作 (操作有三种),
//      将其泛化就是将任意长度为 i 的 word1 转换成任意长度为 j 的 word2 所需最少操作.
//      这里的 DP[i][j] 的定义就是这个意思, 所以需要二重循环 i, j.
func minDistance(word1, word2 string) int {
	return 0
}

func minTriple(p1, p2, p3 int) int {
	if p3 < p2 {
		p2, p3 = p3, p2
	}
	if p1 < p2 {
		return p1
	}
	return p2
}
