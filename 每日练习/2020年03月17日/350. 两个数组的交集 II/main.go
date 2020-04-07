package main

import (
	"fmt"
	"math"
)

func intersect(nums1 []int, nums2 []int) []int {
	mapLenght := int(math.Max(float64(len(nums1)), float64(len(nums2))))
	resultSlice := []int{}

	m := make(map[int]int, mapLenght)
	for _, value := range nums1 {
		if _, ok := m[value]; !ok {
			m[value] = 1
		} else {
			m[value]++

		}
	}
	fmt.Println(m)
	for _, value := range nums2 {
		if _, ok := m[value]; ok {
			if m[value] > 0 {
				resultSlice = append(resultSlice, value)
			}

			m[value]--
			fmt.Println(m)
		}
	}
	fmt.Println(resultSlice)
	return resultSlice
}


func main()  {
	intersect([]int{4,9,5}, []int{9,4,9,8,4})
}