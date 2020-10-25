package leetcode_go_third

type MinStack struct {
	data    []int
	minData []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data:    make([]int, 0),
		minData: make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.minData) == 0 || x < this.minData[len(this.minData)-1] {
		this.minData = append(this.minData, x)
	} else {
		this.minData = append(this.minData, this.minData[len(this.minData)-1])
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.minData = this.minData[:len(this.minData)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minData[len(this.minData)-1]
}
