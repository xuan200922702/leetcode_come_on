package main

import (
	//"fmt"
	//"sort"
	"fmt"
	"sort"
)

//方法1
//
func majorityElement(nums []int) int {
	sort.Ints(nums)
	fmt.Println(nums[int(len(nums)/2)])
	return nums[int(len(nums)/2)]
}



func main(){

	majorityElement([]int{3,2,3,4,5,6,3,5,7,7,5,})
}



