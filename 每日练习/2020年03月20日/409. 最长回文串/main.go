package main

import "fmt"

func longestPalindrome(s string) int {
	var dict [123]int
	for _, v := range s {
		dict[v] += 1
	}
	count := 0
	for _, v := range dict {
		count += v / 2 * 2
	}
	if count == len(s) {
		fmt.Println(count)
		return count
	} else {
		fmt.Println(count+1)
		return count + 1
	}


}

func main()  {
	longestPalindrome("abccccdd")
}