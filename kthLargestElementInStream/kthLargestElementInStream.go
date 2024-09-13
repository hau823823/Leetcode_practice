package kthLargestElementInStream

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 降冪
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type KthLargest struct {
	minHeap *IntHeap
	k       int
}

func Constructor(k int, nums []int) KthLargest {
	tmp := IntHeap(nums)
	this := KthLargest{minHeap: &tmp, k: k}
	heap.Init(this.minHeap)
	for this.minHeap.Len() > k {
		heap.Pop(this.minHeap)
	}

	return this
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.minHeap, val)
	if this.minHeap.Len() > this.k {
		heap.Pop(this.minHeap)
	}
	return (*this.minHeap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
