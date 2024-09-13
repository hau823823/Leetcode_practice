package romanToInteger

func RomanToInt(s string) int {

	roman2IntMap := make(map[string]int, 7)
	roman2IntMap["I"] = 1
	roman2IntMap["V"] = 5
	roman2IntMap["X"] = 10
	roman2IntMap["L"] = 50
	roman2IntMap["C"] = 100
	roman2IntMap["D"] = 500
	roman2IntMap["M"] = 1000

	b := []byte(s)
	l := len(b)
	r := 0
	for i := 0; i < l; i++ {
		if i < l - 1 {
			if roman2IntMap[string(b[i])] < roman2IntMap[string(b[i+1])] {
				r += roman2IntMap[string(b[i+1])]
				r -= roman2IntMap[string(b[i])]
				i ++
				continue
			}
		}
		r += roman2IntMap[string(b[i])]
	}

	return r
}