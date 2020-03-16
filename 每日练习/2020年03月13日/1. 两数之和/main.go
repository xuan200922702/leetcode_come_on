package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for  i:=0;i<len(nums);i++{
		for  j:=1;j<len(nums);j++{
			if nums[i]+nums[j]==target &&i!=j{
				fmt.Println([]int{i,j})
				return  []int{i,j}
			}
		}

	}
	return nil
}
//func twoSum(nums []int, target int) []int {
//	m:=map[int]int{}
//	for i,v :=range(nums){
//		if  k,ok:=m[target-v];ok{
//			return []int{k,i}
//		}
//		m[v]=i
//	}
//	return nil
//}

func main(){
	twoSum([]int{2,7,11,15},9)
}