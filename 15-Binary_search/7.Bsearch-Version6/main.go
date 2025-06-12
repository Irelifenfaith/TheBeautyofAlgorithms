package main

import "fmt"

func bSearch(nums []int, value int) []int {
	//var mid, low, high int
	res := []int{-1, -1}
	n := len(nums)
	fmt.Println(n)
	slow := 0
	high := n - 1
	var mid int
	//找出第一次出现的位置
	for slow <= high {
		// mid = (slow + high) / 2
		// mid = slow + (high-slow)/2
		mid = slow + ((high - slow) >> 2)
		if nums[mid] == value {
			if mid == 0 || nums[mid-1] != value {
				res[0] = mid
				break
			} else {
				high = mid - 1
			}
		} else if nums[mid] > value {
			high = mid - 1
		} else {
			slow = mid + 1
		}

	}
	slow = 0
	high = n - 1
	//找出最后一次出现的位置
	for slow <= high {
		// mid = (slow + high) / 2
		// mid = slow + (high-slow)/2
		mid = slow + ((high - slow) >> 2)
		if nums[mid] == value {
			if mid == high || nums[mid+1] != value {
				res[1] = mid
				break
			} else {
				slow = mid + 1
			}
		} else if nums[mid] > value {
			high = mid - 1
		} else {
			slow = mid + 1
		}

	}

	return res
}

// 二分查找-在排序数组中查找元素的第一个和最后一个位置
// 给你一个按照非递减顺序排列的整数数组 nums，
// 和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
//如果数组中不存在目标值 target，返回 [-1, -1]。
//你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
func main() {
	var value int = 2
	nums := []int{1, 1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, val := range nums {
		fmt.Println(index, val)
	}
	indexDst := bSearch(nums, value)

	fmt.Println(indexDst, "ending.....")

}
