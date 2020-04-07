package main

import "fmt"

func Insertsort(nums [7]int){

	for i:=1;i<len(nums);i++{
	insertVal:=nums[i]
	insertIndex:=i-1

	for  insertIndex >=0 && nums[insertIndex]<insertVal{
		nums[insertIndex+1]=nums[insertIndex]
		insertIndex--
	}
	//插入
	if  insertIndex+1!=i{
		nums[insertIndex+1]=insertVal
	}
		fmt.Println(nums)
	}

}



func main(){
	nums:=[7]int{4,7,1,5,200,152,100}
	Insertsort(nums)
}