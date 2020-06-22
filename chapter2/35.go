package chapter2

// 69 Sqrt(x)
type Sqrter interface {
	MySqrt(x, accuracy int) float64
}

// 1. 二分法
type BinarySqrt struct {
}

// 规律: 误差判断 -> m - x/m
func (bs *BinarySqrt) MySqrt(x float64, accuracy int) float64 {
	acc := bs.accuracy(accuracy)
	var lower, upper float64 = 0, x
	var m float64

	for lower <= upper {
		m = (lower + upper) / 2
		// 计算误差
		x0 := m - x/m
		if x0 < 0 {
			x0 = x/m - m
		}
		if x0 < acc {
			break
		}

		if m < x/m {
			lower = m
		} else {
			upper = m
		}
	}
	return m
}

func (bs *BinarySqrt) accuracy(acc int) float64 {
	var result float64 = 1
	for i := 0; i < acc; i++ {
		result /= 10
	}
	return result
}

// 2. 牛顿迭代
