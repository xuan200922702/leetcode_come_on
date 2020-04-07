package  main

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

//增加一个数到队列
func(this *Queue)Put(value int)(err error){
	if this.isFull(){
		return errors.New("queue is full")
	}
	this.nums[this.tail]=value
	this.tail++
	return
}


//从队列 取出一个数
func(this *Queue)Pop()(value int,err error){
	if this.isNull(){
		return  -1,errors.New("queue is Null")
	}
	value=this.nums[this.head]
	this.head++
	return value,err
}

//展示队列
func(this *Queue)Show(){
	fmt.Println("队列当前的情况是:")
	if this.Size()==0{
		fmt.Println("队列为空")
	}

	temphead:=this.head
	for i:=0;i<this.Size();i++{
		fmt.Printf("nums[%d]==%d\n",temphead,this.nums[temphead])
		temphead=(temphead+1)%this.Maxsize
	}
}

//判断队列是否为空
func(this *Queue)isNull()  bool{
	return  this.tail==this.head
}
//判断队列是否满
func(this *Queue)isFull()bool{
	return (this.tail+1)%this.Maxsize==this.head
}
//判断队列还有多少元素
func(this *Queue)Size()int{
	return (this.tail+this.Maxsize-this.head)%this.Maxsize
}


func main(){
	queue:=&Queue{
		Maxsize: 5,
		head:    0,
		tail:    0,
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
