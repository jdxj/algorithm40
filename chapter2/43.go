package chapter2

// 动态规划 (上)
// 1. 递归+记忆 -> 递推
// 2. 状态的定义: opt[n], dp[n], fib[n]
// 3. 状态转移方程: opt[n] = best_of(opt[n-1], opt[n-2], ...)
// 4. 最优子结构

// 斐波那契数列
//   递推公式: F(n) = F(n-1) + F(n-2)
//   O(2^n)
//              mem: 记忆
func fib(n int, mem map[int]int) int { // O(n)
	if n <= 1 {
		return n
	}
	if _, ok := mem[n]; !ok {
		mem[n] = fib(n-1, mem) + fib(n-2, mem) // 递归
	}
	return mem[n]
}

// 递推化:
func fib2(n int) int {
	if n <= 1 {
		return n
	}

	f := make([]int, n+1)
	f[0], f[1] = 0, 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
