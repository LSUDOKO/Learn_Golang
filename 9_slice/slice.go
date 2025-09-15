package main

import (
	"fmt"
	"slices"
)

func main() {
	// uninitialised slice in nill
	// var nums[]int
	// fmt.Println()
	// fmt.Println(nums==nil)
	// var nums=make([]int,2,5)
	// // capacity -> maximum size of slice
	// fmt.Println(cap(nums))// print 5
	// var nums2=make([]int,0)
	// fmt.Println(nums==nil)
	// fmt.Println(nums2==nil)
	// append
	// var nums=make([]int,2,5)//intial zero element two zero
	// var nums=make([]int,0)// intial no zero element will come
	// nums = append (nums,1)
	// nums = append(nums,2)
	// nums = append(nums,3)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// copy function
	var nums = make([]int, 3, 5)
	// var nums2=make([]int,len(nums),len(nums))
	nums = append(nums, 1, 2, 3, 4, 5)
	// copy(nums2,nums)
	fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(cap(nums2))
	// fmt.Println(nums2)

	// slice
	// nums:=[]int{1,2,3,4}
	// fmt.Println(nums[0:2])

	// slice
	var num1 = []int{1, 2, 3}
	var num2 = []int{1, 2, 3}
	fmt.Println(slices.Equal(num1, num2))
}
