package main

import (
	"fmt"
	"leetcode/questions/low"
)

func main(){

	obj := low.Constructor()
	obj.AppendTail(1)
	obj.AppendTail(8)
	obj.AppendTail(20)
	obj.AppendTail(1)
	obj.AppendTail(11)
	obj.AppendTail(2)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}
