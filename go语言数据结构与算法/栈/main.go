package main

import (
	"errors"
	"fmt"
)

type  stack struct {
	MaxTop int
	Top   int
	nums  [5]int
}
//入栈
func (this *stack)Push(val  int)(err error)  {
	//先判断是否满了
	if  this.Top ==this.MaxTop{
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.nums[this.Top]=val
	return
}

//出栈
func(this *stack)Pop()(val int,err  error){
	if  this.Top==-1{
		fmt.Println("stack is  empty")
	}
	val=this.nums[this.Top]
	this.Top--
	return val,nil
}


//遍历栈
func(this  *stack)List(){
	if this.Top==-1{
		fmt.Println("stack empty")
		return
	}
	for i:=this.Top;i>=0;i--{
		fmt.Printf("nums[%d]=%d\n",i,this.nums[i])
	}
}

func main()  {
	stack:=&stack{
		MaxTop: 5,
		Top:   -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	//显示
	stack.List()
	fmt.Println()
	val,_  := stack.Pop()
	val,_  = stack.Pop()
	val,_  = stack.Pop()
	val,_  = stack.Pop()
	val,_  = stack.Pop()
	fmt.Println("出栈val=",val)
	stack.List()
}