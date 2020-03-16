package main

import "fmt"

func  maxProfit(princes []int) int{
	if len(princes)<2{
		return 0
	}
	prince:=0
	for i:=0;i<len(princes)-1;i++{
		if  princes[i]<princes[i+1]{
			prince+=princes[i+1]-princes[i]
		}
	}
	fmt.Println(prince)
	return prince
}


//第一次自己做
//func maxProfit(prices []int) int {
//	if len(prices)<2{
//		return 0
//	}
//    d:=make([]int,len(prices))
//	price:=0
//	for i:=0;i<len(prices);i++{
//		for j:=1;j<len(prices);j++{
//			cur:=0
//			if prices[j]>prices[i]{
//				cur=cur+prices[j]-prices[i]
//				d[i]=cur
//			}else{
//				d[i]=0
//			}
//			price=max(price,d[i])
//	}
//
//	}
//
//	fmt.Println(price)
//	return price
//}
func max(a,b int)int  {
	if a>b{
		return a
	}

	return b
}


func  main(){
	maxProfit([]int{7,1,5,3,6,4})
}
