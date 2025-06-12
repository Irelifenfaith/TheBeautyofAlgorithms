package main

import "fmt"

func bSearch(nums []int, value int) int {
	//var mid, low, high int
	n := len(nums)
	fmt.Println(n)
	slow := 0
	high := n - 1
	var mid int
	for slow <= high {
		// mid = (slow + high)/2
		// mid = slow+(high-slow)/2
		mid = slow + ((high - slow) >> 1)
		if nums[mid] == value {
			return mid
		} else if nums[mid] > value {
			high = mid - 1
		} else {
			slow = mid + 1
		}
	}
	return -1
}

// 二分查找
func main() {
	var value int = 2
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, val := range nums {
		fmt.Println(index, val)
	}
	indexDst := bSearch(nums, value)

	fmt.Println(indexDst, "ending.....")

}
