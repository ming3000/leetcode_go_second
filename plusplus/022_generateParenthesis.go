package plusplus

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	dfsGenerateParenthesis(0, 0, n, "", &ret)
	return ret
}

func dfsGenerateParenthesis(leftCount, rightCount int, n int, path string, ret *[]string) {
	if leftCount == n && rightCount == n {
		temp := path
		*ret = append(*ret, temp)
		return
	}
	if leftCount < n {
		dfsGenerateParenthesis(leftCount+1, rightCount, n, path+"(", ret)
	}
	if rightCount < leftCount && rightCount < n {
		dfsGenerateParenthesis(leftCount, rightCount+1, n, path+")", ret)
	}
}
