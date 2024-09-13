package countPrimes

// sieve of eratosthenes
// 思路：從 2 開始遍歷到根號 n，然後將所有質數以及其倍數標記出來，剩下未被標記的就是質數
func CountPrimes(n int) int {
	if n < 2 {
		return 0
	}

	notPrime := make([]bool, n)
	count := 0
	
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count ++
			for j := 2; i*j < n; j++ {
				notPrime[i*j] = true
			}
		}
	}

	return count
}