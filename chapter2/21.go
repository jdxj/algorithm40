package chapter2

// 递归
// 分治

// 递归模板
const limit = 10

func recursion(param1 int, param2 string) {
	// part1: 终止条件
	if param1 > limit {
		return
	}

	// part2: 进入下一层处理
	recursion(param1+1, param2)

	// part3: 收尾/返回需要的值
	return // some value
}

// 分治模板
func divideConquer(param1 int, param2 string) string {
	// part1: 终止条件
	if param1 > limit {
		return ""
	}

	// part2: 进入下一层处理
	subresult1 := divideConquer(param1+1, param2)
	subresult2 := divideConquer(param1+1, param2)
	subresult3 := divideConquer(param1+1, param2)

	// part3: 合并数据
	result := subresult1 + subresult2 + subresult3
	return result // some value
}
