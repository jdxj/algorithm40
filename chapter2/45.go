package chapter2

// 70 爬楼梯

// 1. 回溯 (规律: 斐波那契)
//                            f(n): 到第 n 阶, 有多少种
//                           ----
//                  f(n-1) |
//                 -------
//        f(n-2) |
//       -------
//  o  |
//   因为想要到达最顶层 (fn) 之前, 要么使用1步 (f(n-1)), 要么使用2步 (f(n-2)),
//   所以, 到达 f(n) 的走法 = f(n-1) + f(n-2).

// 2. 动态规划
//   不使用数组保存中间变量
func fib3(n int) int {
	if n <= 1 {
		return n
	}

	before2, before1 := 0, 1
	for i := 2; i <= n; i++ {
		before2, before1 = before1, before2+before1
	}
	return before1
}
