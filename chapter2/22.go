package chapter2

// 50 实现 pow(x, y)
type Pow interface {
	MyPow(x float64, n int) float64
}

// 1. 库函数 O(1)
// 2. 暴力循环 O(n)

// 3. 分治
// 原理类似: 计算 2^8 = 2*2*2*2*2*2*2*2
//   2*2   = 4    2 * 2              2次
//   4*4   = 16   2*2 * 2*2          4次
//   16*16 = 256  2*2*2*2 * 2*2*2*2  8次
//   总共需要3次计算
type DivideConquerPow struct {
}

func (dc *DivideConquerPow) MyPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 { // 如果 n 小于0, 那么先计算分母: -n
		return 1 / dc.MyPow(x, -n)
	}
	if n%2 != 0 { // 如果 n 是奇数, 那么变成偶数
		return x * dc.MyPow(x, n-1)
	}
	// n 最小为2
	// 第一个参数可以理解为: x 已乘了多少次, 变化是已乘了: 2次, 4次, 8次
	// 第二个参数可以解释为还需翻倍次数
	return dc.MyPow(x*x, n/2)
	// 2*2   4
	// 4*4   2
	// 16*16 1
	// 16*16 0
}

// x^y = (x*x)^(logy)
// 2*8 = (2*2)^(log8) = (2*2)^3, 只算3次
// 如果 n 不为负, pow 可以简化为:
func myPow(x float64, n uint) float64 {
	if n == 0 {
		return 1
	}
	if n%2 != 0 {
		return x * myPow(x, n-1)
	}
	return myPow(x*x, n/2)
}

// 迭代
type IteratePow struct {
}

func (ip *IteratePow) MyPow(x float64, n int) float64 {
	if n < 0 {
		x = 1 / x
		n = -n
	}

	var pow float64 = 1
	for ; n != 0; n >>= 1 {
		if n&1 != 0 { // 奇数
			pow *= x
		}
		x *= x
	}
	return pow
}
