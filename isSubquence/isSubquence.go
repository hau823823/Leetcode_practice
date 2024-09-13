package isSubquence

// two pointer
func IsSubsequence(s string, t string) bool {
	ptrS, ptrT := 0, 0

	for ptrS < len(s) && ptrT < len(t) {
		if s[ptrS] == t[ptrT] {
			ptrS++
		}
		ptrT++
	}

	return ptrS == len(s)
}
