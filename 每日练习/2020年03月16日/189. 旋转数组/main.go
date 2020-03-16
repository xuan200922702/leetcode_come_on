package main

func rotate(nums []int, k int)  {
	k %= len(nums)
	nums1:=append(nums[len(nums)-k:],nums[:len(nums)-k]...)
	nums=append(nums[:0],nums1...)

}


//自己做的
//func rotate(nums []int, k int)  {
//	nums2:=[]int{}
//	nums3:=[]int{}
//	for i:=0;i<len(nums);i++{
//		if i>k {
//			nums2 = append(nums2, nums[i])
//
//		}else {
//			nums3 = append(nums3, nums[i])
//		}
//
//	}
//	 nums=append(nums2,nums3...)
//     fmt.Println(nums)
//}

func main(){
	rotate([]int{1,2,3,4,5,6,7},3)
}