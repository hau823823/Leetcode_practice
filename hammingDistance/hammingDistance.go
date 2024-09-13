package hammingDistance

func HammingDistance(x int, y int) int {
	compare := uint32(x) ^ uint32(y)
	r := 0

	for {
		if compare == 0 {
			break
		}
		if compare&1 == 1{
			r ++
		}
		compare = compare >> 1
	}

	return r
}