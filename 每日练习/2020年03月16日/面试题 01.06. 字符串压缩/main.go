package main

import (
	"fmt"
	"strconv"
)
//双指针解法
func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}
	res:=""
	slow,quick :=0,0
	for quick <len(S){
		if  S[quick]!=S[slow]{
			res+=string(S[slow])+strconv.Itoa(quick-slow)
			slow=quick
		}
	quick++
	}
	res +=string(S[slow]) + strconv.Itoa(quick-slow)
	if len(res) >= len(S) {
		return S
	}
	fmt.Println(res)
	return res
}

func  main()  {
	compressString("aabcccccaaa")
}

