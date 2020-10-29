package middle

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	dfsGenerateParenthesis(0, 0, n, "", &res)
	return res
}

func dfsGenerateParenthesis(lc, rc int, n int, path string, res *[]string) {
	if lc == n && rc == n {
		*res = append(*res, path)
	}

	if lc < n {
		dfsGenerateParenthesis(lc+1, rc, n, path+"(", res)
	}
	if rc < n && rc < lc {
		dfsGenerateParenthesis(lc, rc+1, n, path+")", res)
	}
}
