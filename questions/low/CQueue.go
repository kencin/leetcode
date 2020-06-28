//剑指 Offer 09. 用两个栈实现队列
// 2020.06.28 22:29 by Kencin

package low

type CQueue struct {
	stack [2][] int
	currentStack int
}


func Constructor() CQueue {
	return CQueue{currentStack: 0}
}


func (this *CQueue) AppendTail(value int)  {
	this.push(this.currentStack, value)
}


func (this *CQueue) DeleteHead() int {
	if !this.isEmpty(1){
		return this.pop(1)
	}
	for !this.isEmpty(this.currentStack){
		this.push(1, this.pop(this.currentStack))
	}
	return this.pop(1)
}

func (this *CQueue) push(stackNum int, value int){
	this.stack[stackNum] = append(this.stack[stackNum], value)
}

func (this *CQueue) pop(stackNum int) int{
	currentStackSize := len(this.stack[stackNum])
	if currentStackSize == 0{
		return -1
	}
	value := this.stack[stackNum][currentStackSize - 1]
	this.stack[stackNum] = this.stack[stackNum][:currentStackSize - 1]
	return value
}

func (this *CQueue) isEmpty(stackNum int) bool {
	return len(this.stack[stackNum]) == 0
}

func (this *CQueue) peek(stackNum int) int {
	return this.stack[stackNum][len(this.stack[stackNum]) - 1]
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */