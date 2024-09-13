package generateParenthesis

func GenerateParenthesis(n int) []string {
    result := make([]string, 0)

	var backtrack func(string, int, int)
	backtrack = func(str string, open, close int) {
		if len(str) == 2 * n {
			result = append(result, str)
		}

		if open < n {
			backtrack(str+"(", open + 1, close)
		}

		if close < open {
			backtrack(str+")", open, close + 1)
		}
	}

	backtrack("", 0, 0)
	return result
}