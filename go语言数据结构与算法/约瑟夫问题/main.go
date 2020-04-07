package main

import (
	"fmt"
)
//分析思路
//编写一个函数 ，playGame(first *Boy,starNo int)
//留下最后一个人
type Boy struct {
	No int
	Next *Boy
}

func AddBoy(num int) *Boy {
	first:=&Boy{}  //空节点
	curBoy  :=&Boy{}
	if num<1{
		fmt.Println("num的值不对")
	}

	for i:=1;i<=num;i++{
		boy:=&Boy{
			No:   i,
			Next: nil,
		}
		if i==1{
			first=boy
			curBoy=boy
			curBoy.Next=first
		}else {
			curBoy.Next=boy
			curBoy=boy
			curBoy.Next=first
		}

	}
	return  first
}

func showBoy(first *Boy)  {
	if first.Next==nil{
		fmt.Println("链表 为空，没有小孩")
		return
	}
	curBoy:=first
	for {
		fmt.Printf("小孩编号=%d ->",curBoy.No)
		if  curBoy.Next==first{
			break
		}
		curBoy=curBoy.Next
	}
}


func PlayGame(first *Boy,startNo int,countNum int){
	if   first.Next==nil{
		fmt.Println("空的链表")
		return
	}
	tail:=first
	//tail到了最后
	for {
		if tail.Next==first{
			break
		}
		tail=tail.Next
	}
	//让first移动到startNo
	for  i:=1;i<=startNo-1;i++{
		first=first.Next
		tail=tail.Next

	}
	for{
		for i:=1;i<=countNum-1;i++{
			first=first.Next
			tail=tail.Next
		}
		fmt.Printf("小孩编号为%d 出圈\n",first.No)
		//删除first执行的小孩
		first=first.Next
		tail.Next=first
		//判断
		if tail==first{
			break
		}
	}
	fmt.Printf("小孩编号为%d 出圈\n",first.No)

}

func main()  {
	first:=AddBoy(5)

	PlayGame(first,2,3)
}