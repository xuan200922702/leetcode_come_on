package main

import (
	"fmt"
)

type HeroNode struct {
	no  int
	name  string
	nickname string
	next   *HeroNode

}

//在单链表 后插入
//编写 第二种插入方法，根据no的编号从小到大插入
func InsertHeroNode2(head *HeroNode,newHeroNode *HeroNode) {
	//先找到该链表的最后这个节点
	//创建 一个辅助 节点
	temp := head
	flag:=true
	for {
		if temp.next == nil {
			break
		}else if temp.next.no  >= newHeroNode.no{
			break
		}else if temp.next.no == newHeroNode.no{
			flag=false
			break
		}
		temp=temp.next
	}
	if !flag{
		fmt.Println("对不起，已经存在",newHeroNode.no)
		return
	}else {
		newHeroNode.next=temp.next
		fmt.Println(newHeroNode.next)
		temp.next=newHeroNode

		return
	}

}

func  DeleteHerNode(head  *HeroNode,id  int){
	temp:=head
	flag:=false
	for {
		if temp.next ==nil{
			break
		}else if  temp.next.no== id {
			flag=true
			break
		}
		temp=temp.next
	}
	if  flag{
		temp.next=temp.next.next
	}else {
		fmt.Println("要删除的id不存在")
	}
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
	//hero4:=&HeroNode{
	//	no:3,
	//	name:"吴用",
	//	nickname:"智多星",
	//}

	InsertHeroNode2(head,hero1)
	InsertHeroNode2(head,hero3)
	InsertHeroNode2(head,hero2)
	//InsertHeroNode2(head,hero4)
	ListHeroNode(head)
	fmt.Println()
	DeleteHerNode(head,2)

	ListHeroNode(head)
}
