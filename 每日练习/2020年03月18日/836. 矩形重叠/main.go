package main

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	if  rec2[0]>=rec1[2] || rec2[2]<=rec1[0]||rec2[1]>=rec1[3]||
		rec2[3]<=rec1[1]{
		return false
	}
	return true


}


func main(){
	isRectangleOverlap([]int{0,0,2,2},[]int{1,1,3,3})
}