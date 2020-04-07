package main

import "fmt"

func Selectsort(nums [7]int){

	for j:=0;j<len(nums)-1;j++{
	maxnum:=nums[j]
	maxindex:=j
	for i:=j+1;i<len(nums);i++{
		if maxnum<nums[i]{
			maxnum=nums[i]
			maxindex=i
		}
	}
	if maxindex!=j{
		nums[j],nums[maxindex]=nums[maxindex],nums[j]
	}
	}
	fmt.Println(nums)
}



func main(){
  nums:=[7]int{4,7,1,5,200,152,100}
  Selectsort(nums)
}