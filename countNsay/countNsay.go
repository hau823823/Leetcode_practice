package countNsay

func CountAndSay(n int) string {
	res := "1"

	for i := 0; i < n-1; i++ {
		tmp := Helper1(res)
		res = Helper2(tmp)
	}

	return res
}

func Helper1(str string) [][]int {
	ret := make([][]int, 0)
	tmp := rune(str[0])
	count := 0

	for _, s := range str {
		if s == tmp {
			count++
		} else {
			ret = append(ret, []int{count, int(tmp - '0')})
			tmp = s
			count = 1
		}
	}
	ret = append(ret, []int{count, int(tmp - '0')})

	return ret
}

func Helper2(record [][]int) string {
	tmp := []byte{}
	for _, r := range record {
		for _, v := range r {
			tmp = append(tmp, byte(v+'0'))
		}
	}
	return string(tmp)
}
