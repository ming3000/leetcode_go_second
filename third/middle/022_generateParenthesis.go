package middle

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfsGenerateParenthesis(0, 0, n, "", &res)
	return res
}

func dfsGenerateParenthesis(leftCount, rightCount int, n int, path string, res *[]string) {
	if leftCount == n && rightCount == n {
		*res = append(*res, path)
	}
	if leftCount < n {
		dfsGenerateParenthesis(leftCount+1, rightCount, n, path+"(", res)
	}
	if rightCount < n && rightCount < leftCount {
		dfsGenerateParenthesis(leftCount, rightCount+1, n, path+")", res)
	}
}
