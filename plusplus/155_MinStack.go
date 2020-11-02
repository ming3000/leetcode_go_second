package plusplus

// 我们使用两个栈：
// 一个栈存放全部的元素，push，pop都是正常操作这个正常栈。
// 另一个存放最小栈。
// 每次push，如果比最小栈的栈顶还小，我们就push进最小栈，否则不操作
// 每次pop的时候，我们都判断其是否和最小栈栈顶元素相同，如果相同，那么我们pop掉最小栈的栈顶元素即可
//
// 关键点
// 往minstack中 push的判断条件。 应该是stack为空或者x小于等于minstack栈顶元素
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
	this.minData = this.data[:len(this.minData)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.minData[len(this.minData)-1]
}
