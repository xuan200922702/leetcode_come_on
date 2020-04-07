package main

import (
	"fmt"
)
func moveZeroes(nums []int)  {
	l:=len(nums)
	i:=0
	for {
		if i>=l{
			break
		}
		if nums[i]==0 {
			l--
			nums = append(nums[:i], nums[i+1:]...)
			nums=append(nums,0)
		}else {
			i++
		}
	}

	fmt.Println(nums)
}

func  main(){
	moveZeroes([]int{0,1,0,3,12})
}