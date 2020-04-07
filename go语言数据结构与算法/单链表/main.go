package main

import "fmt"

type HeroNode struct {
	no  int
	name  string
	nickname string
	next   *HeroNode

}

//在单链表 后插入
//编写 第二种插入方法，根据no的编号从小到大插入
func InsertHeroNode(head *HeroNode,newHeroNode *HeroNode){
	//先找到该链表的最后这个节点
	//创建 一个辅助 节点
	temp:=head
	for{
		if temp.next ==nil{
			break
		}
		temp=temp.next
	}
	//将newHeroNode加入链表最后
	temp.next=newHeroNode
}

//显示 链表的 所有 信息

func  ListHeroNode(head *HeroNode){
	temp:=head
	//先判断该链表是不是为空链表
	if temp.next==nil{
		fmt.Println("该链表为 空")
		return
	}
	//遍历这个链表
	for{
		fmt.Printf("[%d,%s,%s]==>",temp.next.no,temp.next.name,temp.next.nickname)
		//判断是否到最后
		temp=temp.next
		if temp.next==nil{
			break
		}
	}

}

func main()  {
	//创建一个头节点
	head :=&HeroNode{}
	//创建一个新的节点
	hero1:=&HeroNode{
		no:1,
		name:"宋江",
		nickname:"及时雨",
	}
	hero2:=&HeroNode{
		no:2,
		name:"卢俊义",
		nickname:"玉麒麟",
	}
	hero3:=&HeroNode{
		no:3,
		name:"林冲",
		nickname:"豹子头",
	}
	InsertHeroNode(head,hero3)
	InsertHeroNode(head,hero1)
	InsertHeroNode(head,hero2)
	ListHeroNode(head)
}