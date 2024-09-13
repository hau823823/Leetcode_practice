package find1StOccurenceInStr

func StrStr(haystack string, needle string) int {
    pt1 := 0
	nl := len(needle)

	for pt1 < len(haystack) - nl + 1 {
		if haystack[pt1:pt1+nl] == needle{
			return pt1
		}
		pt1++
	}

	return -1
}