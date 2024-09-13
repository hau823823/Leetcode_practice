package shuffleArrary

import "math/rand"

type Solution struct {
	array []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

func (this *Solution) Reset() []int {
	return this.array
}

func (this *Solution) Shuffle() []int {
	l := len(this.array)
	res := make([]int, l)
	copy(res, this.array)
	for i := 0; i < l; i++ {
		j := rand.Intn(l - i)
		res[i], res[j + i] = res[j + i], res[i]
	}
	return res
}
