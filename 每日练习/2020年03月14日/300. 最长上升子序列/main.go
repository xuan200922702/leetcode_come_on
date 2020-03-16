package main

func lengthOfLIS(nums []int) int {
	if nums==nil || len(nums)==0{
		return 0
	}
	d:=make([]int,len(nums))
	d[0]=1
	m:=1
	for i:=1;i<len(nums);i++{
		for j:=0;j<i;j++{
			cur:=1
			if nums[j]<nums[i]{
				cur=d[j]+1
			}
			d[i]=max(cur,d[i])
		}
		m=max(m,d[i])
	}
	return m
}

func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}
func main(){
	lengthOfLIS([]int{10,9,2,5,3,7,101,18})
}