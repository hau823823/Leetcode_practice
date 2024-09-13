package letterCombinationsOfPhoneNumber

func LetterCombinations(digits string) []string {
	set := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := make([]string, 0)

	var backtrack func(int, string)
	backtrack = func(pos int, s string) {
		if pos == len(digits) {
			result = append(result, s)
			return
		}

		for _, c := range set[digits[pos]] {
			backtrack(pos+1, s + string(c))
		}
	}

	if len(digits) > 0 {
        backtrack(0, "")
    }
	return result
}