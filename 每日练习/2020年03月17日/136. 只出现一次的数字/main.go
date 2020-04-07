package main

import "fmt"

func singleNumber(nums []int) int {
	haha:=make(map[int]int)
	for _,v :=range nums{
		if  _,ok:=haha[v];ok{
			delete(haha,v)
		}else{
			haha[v]=0
		}
	}
	for  i:=range haha{
		fmt.Println(i)
		return i
	}

	return 0
}

func main()  {
	singleNumber([]int{4,1,2,1,2})
}