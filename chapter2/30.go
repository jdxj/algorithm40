package chapter2

// 22 括号生成
type BracketGenerator interface {
	GenerateParenthesis(n int) []string
}

// 1. 数学归纳

// 2. 递归 (DFS)
//   a. 枚举所有情况
//   b. 验证

// 3. 改进方法2
//   - 生成时按照一定规则
// 规律是一个一个往后添加 "(/)"
type RecursionBG struct {
	result []string
	limit  int
}

func (rbg *RecursionBG) GenerateParenthesis(n int) []string {
	rbg.limit = n
	rbg.helper(0, 0, "")
	return rbg.result
}

func (rbg *RecursionBG) helper(left, right int, result string) {
	// 次数用完, 则保存结果
	if left == rbg.limit && right == rbg.limit {
		rbg.result = append(rbg.result, result)
		return
	}

	// 次数没用完就继续添加括号
	// 尝试添加左括号
	if left < rbg.limit {
		rbg.helper(left+1, right, result+"(")
	}
	// 尝试添加右括号
	if right < rbg.limit && left > right {
		rbg.helper(left, right+1, result+")")
	}
}
