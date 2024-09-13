package minStack

type MinStack struct {
    arrary []int
	min    []int
}


func Constructor() MinStack {
    return MinStack{[]int{}, []int{}}
}


func (this *MinStack) Push(val int)  {
    this.arrary = append(this.arrary, val)

	if len(this.min) == 0 {
		this.min = append(this.min, 0)
	} else if this.arrary[this.min[len(this.min) - 1]] < val {
		this.min = append(this.min, this.min[len(this.min) - 1])
	} else {
		this.min = append(this.min, len(this.arrary) - 1)
	}
}


func (this *MinStack) Pop()  {
    this.arrary = this.arrary[:len(this.arrary) - 1]
	this.min = this.min[:len(this.min) - 1]
}


func (this *MinStack) Top() int {
    l := len(this.arrary)
	return this.arrary[l-1]
}


func (this *MinStack) GetMin() int {
    return this.arrary[this.min[len(this.min) - 1]]
}
