// 面试题 03.01. 三合一
// 2020.06.28 21:20 by kencin

// 2020.06.28 21.48 Done

package low

type TripleInOne struct {
	everyStackSize int
	stack [3][] int
}


func constructor(stackSize int) TripleInOne {
	if stackSize < 0{
		panic("stack size can not < 0")
	}
	return TripleInOne{everyStackSize: stackSize}
	//var tripleInOne TripleInOne
	//tripleInOne.stackSize = stackSize
	//return tripleInOne
}


func (this *TripleInOne) push(stackNum int, value int)  {
	// 栈满，跳过
	if len(this.stack[stackNum]) >= this.everyStackSize {
		return
	}
	// 压入元素
	this.stack[stackNum] = append(this.stack[stackNum], value)
}


func (this *TripleInOne) pop(stackNum int) int {
	currentStackSize := len(this.stack[stackNum])
	if currentStackSize == 0{
		return -1
	}
	value := this.stack[stackNum][currentStackSize - 1]
	this.stack[stackNum] = this.stack[stackNum][:currentStackSize - 1]
	return value
}


func (this *TripleInOne) peek(stackNum int) int {
	currentStackSize := len(this.stack[stackNum])
	if currentStackSize == 0{
		return -1
	}
	return this.stack[stackNum][currentStackSize - 1]
}


func (this *TripleInOne) isEmpty(stackNum int) bool {
	return len(this.stack[stackNum]) == 0
}


//func main(){
//	obj := low.Constructor(2)
//	obj.Push(0, 1)
//	obj.Pop(0)
//	fmt.Println(obj.IsEmpty(0))
//	fmt.Println(obj.Peek(0))
//}