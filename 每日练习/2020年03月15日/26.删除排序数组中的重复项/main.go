package main

func removeDuplicates(nums []int) int {
	for i :=len(nums)-1;i>0;i--{
		if nums[i]==nums[i-1]{
			nums=append(nums[:i],nums[i+1:]...)
		}
	}
	return  len(nums)
}

func main(){
	removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4})
}
