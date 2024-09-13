package fraction2RecurringDecimal

import "strconv"

func FractionToDecimal(numerator int, denominator int) string {
	res := make([]byte, 0)

	if numerator == 0 {
		return "0"
	} else if (numerator < 0) != (denominator < 0) {
		res = append(res, '-')
	}

	num := abs(numerator)
	den := abs(denominator)
	quot := num / den
	res = append(res, []byte(strconv.Itoa(quot))...)

	rMap := make(map[int]int)
	dig := make([]byte, 0)
	remain := num % den
	count := 0
	rMap[remain] = count
	for remain != 0 {
		quot = remain * 10 / den
		remain = remain * 10 % den
		dig = append(dig, []byte(strconv.Itoa(quot))...)

		idx, exists := rMap[remain]
		if exists {
			dig = append(dig[:idx], append([]byte{'('}, dig[idx:]...)...)
			dig = append(dig, ')')
			break
		}

		count++
		rMap[remain] = count
	}

	if len(dig) > 0 {
		res = append(res, '.')
		res = append(res, dig...)
	}

	return string(res)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
