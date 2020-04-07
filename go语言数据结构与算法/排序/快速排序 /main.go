package main

import "fmt"

func QuckSort(left int,right  int,nums  *[7]int){
	l:=left
	r:=right
	pivot:=nums[(left+right)/2]
	temp:=0
	for l<r{
		for nums[l]<pivot{
			l++
		}
		for nums[r]>pivot{
			r--
		}
		if  l>=r{
			break
		}
		temp=nums[l]
		nums[l]=nums[r]
		nums[r]=temp
		if  nums[l]==pivot{
			r--
		}
		if   nums[r]==pivot{
			l++
		}
	}
	if  l==r{
		l++
		r--
	}
	//向左递归
	if left <r{
		QuckSort(left,r,nums)
	}
	//向右递归
	if right>l{
		QuckSort(l,right,nums)
	}

}


func main(){
	nums:=[7]int{4,7,1,5,200,152,100}
	QuckSort(0,len(nums)-1,&nums)
	fmt.Println(nums)
}
