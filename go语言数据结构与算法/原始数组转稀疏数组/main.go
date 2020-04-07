package main

import (
	"fmt"
)

type  ValNode struct {
	row   int
	col   int
	value  int
}

func main()  {
	var chessMap[11][11]int
	chessMap[0][2]=3
	chessMap[3][4]=7

	//输出原始数组
	for  _,v :=range chessMap{
		for _,k:=range v{
			fmt.Printf("%d\t",k)
		}
		fmt.Println()
	}


	//找出对应的
	var spar []ValNode
	for i,v:=range chessMap{
		for k,v2:=range v{
			if  v2!=0{
				valnode:=ValNode{
					row:i,
					col:k,
					value:v2,
				}
				spar=append(spar,valnode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前稀疏数组是 ")
	fmt.Println("row   col  value")
	for i,valnode:=range spar{
		fmt.Printf("%d %d  %d   %d\n",i,valnode.row,valnode.col,valnode.value)
	}
}
