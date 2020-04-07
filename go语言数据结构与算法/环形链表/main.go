package main

import "fmt"

type  CatNode struct {
	no int
	name  string
	next *CatNode
}

func InsertCatNode(head  *CatNode ,newCatNode  *CatNode)  {
	//判读是不是添加第一只猫
	if  head.next == nil{
		head.no=newCatNode.no
		head.name=newCatNode.name
		head.next=head
		fmt.Println(newCatNode,"加入到环形链表 ")
		return
	}
	temp:=head
	for{
		if temp.next==head{
			break
		}
	}
	temp.next=newCatNode
	newCatNode.next=head

}

func DelCaNode(head *CatNode,id int) *CatNode{
	temp:=head
	helper:=head
	flag:=true
	if temp.next   == nil{
		fmt.Println("这是一个空的环形链表")
		return   head
	}
	//如果 只有一个节点
	if temp.next==head{
		temp.next=nil
		return head
	}
	//helper  定位到  链表最后
	for{
		if helper.next  ==head{
			break
		}
		helper = helper.next
	}

	//如果包含两个以上节点
	for{
		if  temp.next==head{
			break
		}
		if temp.no==id{
			if temp==head{
				head=head.next
			}
			helper.next = temp.next
			fmt.Printf("猫猫=%d",id)
			flag=false
			break
		}
		temp=temp.next
		helper=helper.next
	}
	if flag{
		if temp.no==id{
			helper.next=temp.next
			fmt.Printf("猫猫=%d\n",id)
		}else{
			fmt.Println("对不起，没有这个猫猫",id)
		}
	}
	return  head
}




func ListCircleLink(head *CatNode)  {
	fmt.Println("环形链表的情况是")
	temp:=head
	if  temp.next==nil{
		fmt.Println("空的")
		return
	}
	for{
		fmt.Printf("猫的信息[id=%d name=%s]->",temp.no,temp.name,)
		if  temp.next == head{
			break
		}
		temp =temp.next
	}
}

func  main(){
	head:=&CatNode{}

	//创建 一只猫
	cat1:=&CatNode{
		no:   1,
		name: "tome",
	}
	cat2:=&CatNode{
		no:   2,
		name: "tome",
	}
	InsertCatNode(head,cat1)
	InsertCatNode(head,cat2)
	ListCircleLink(head)
	fmt.Println()
	head=DelCaNode(head,1)
	ListCircleLink(head)
}