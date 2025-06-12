package main

import "fmt"

func bSearch(nums []int, value int) int {
	//var mid, low, high int
	n := len(nums)
	fmt.Println(n)
	var slow, high, index int

	// 找到最小值对应的下标
	for i := 1; i < n; i++ {
		if nums[i-1] > nums[i] {
			index = i
			break
		}
	}
	// 目标值和切片中第一个元素和最后一个元素比较
	// 初始化slow和high
	if value <= nums[n-1] {
		slow = index
		high = n - 1
	} else {
		slow = 0
		high = index - 1
	}

	var mid int
	for slow <= high {
		// mid = (slow + high)/2
		// mid = slow+(high-slow)/2
		mid = slow + ((high - slow) >> 1)
		if nums[mid] == value {
			return mid
		} else if nums[mid] < value {
			slow = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

// 二分查找-循环有序数组中查找等于给定值的元素
func main() {
	var value int = 2
	nums := []int{7, 8, 9, 10, 1, 2, 3, 4, 5, 6}
	for index, val := range nums {
		fmt.Println(index, val)
	}
	indexDst := bSearch(nums, value)

	fmt.Println(indexDst, "ending.....")

}
