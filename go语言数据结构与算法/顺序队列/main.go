package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Maxsize   int
	nums      [5]int
	head      int
	tail      int

}

//添加数据队列
func (this *Queue)Put(value int)(err error){
	if this.tail==this.Maxsize-1{
		return  errors.New("queue is full")
	}
	this.tail++
	this.nums[this.tail]=value
	return
}

//从队列取出数据

func (this  *Queue)Pop()(value  int,err error){
	if   this.tail==this.head{
		return  -1, errors.New("queue is null")
	}
	this.head++
	value=this.nums[this.head]
	return value,err
}

//展示队列 
func (this  *Queue)Show(){
	fmt.Println("队列当前的情况是:")
	for i:=this.head+1;i<=this.tail;i++{
		fmt.Printf("nums[%d]=%d\n",i,this.nums[i])
	}
	return
}

func main(){
	queue:=&Queue{
		Maxsize: 5,
		head:    -1,
		tail:    -1,
	}
	var key string
	var  value int
	for{
		fmt.Println("===========")
		fmt.Println("1.show队列")
		fmt.Println("2.put队列")
		fmt.Println("3.pop队列")
		fmt.Println("===========")
		fmt.Scanln(&key)
		switch key {
		case "put":
			fmt.Println("请输入 一个数")
			fmt.Scanln(&value)
			err:=queue.Put(value)
			if   err!=nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println("put到队列成功")
			}
		case "pop":
			value,err:=queue.Pop()
			if   err!=nil{
				fmt.Println(err.Error())
			}else{
				fmt.Printf("pop一个数到队列成功==%d\n",value)
			}
		case "show":
			queue.Show()

		}
	}
}