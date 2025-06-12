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
			// 如果mid等于high，那这个元素已经是数组的最后一个元素，那它肯定是我们要找的；
			// 如果mid不等于high，但nums[mid]的后一个元素nums[mid+1]不等于value，
			// 那也说明nums[mid]就是我们要找的最后一个值等于给定值的元素
			// 如果经过检查之后发现nums[mid]后面的一个元素nums[mid+1]也等于value，
			// 那说明此时的nums[mid]肯定不是我们要查找的最后一个值等于给定值的元素。
			// 那我们就更新slow = mid + 1，因为要找的元素肯定出现在[mid+1, high]之间。
			if mid == high || (nums[mid+1] != value) {
				return mid
			} else {
				slow = mid + 1
			}
		} else if nums[mid] > value {
			high = mid - 1
		} else {
			slow = mid + 1
		}
	}
	return -1
}

// 二分查找-查找最后1个值等于给定值的元素
func main() {
	var value int = 2
	nums := []int{1, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, val := range nums {
		fmt.Println(index, val)
	}
	indexDst := bSearch(nums, value)

	fmt.Println(indexDst, "ending.....")

}
