package ransomNote

func CanConstruct(ransomNote string, magazine string) bool {
    set := make(map[rune]int)
	sett := make(map[rune]int)

	for _, m := range magazine {
		set[m]++
	}

	for _, r :=  range ransomNote {
		if _, exists := set[r]; !exists {
			return false
		}
		sett[r]++
	}

	for _, r :=  range ransomNote {
		if set[r] < sett[r] {
			return false
		}
	}

	return true
}

func CanConstructOp(ransomNote string, magazine string) bool {
    set := [26]int{}

	for _, m := range magazine {
		set[m - 'a']++
	}

	for _, r :=  range ransomNote {
		if set[r - 'a'] == 0 {
			return false
		}
		set[r - 'a']--
	}

	return true
}