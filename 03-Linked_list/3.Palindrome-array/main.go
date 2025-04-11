package main

import "fmt"

// 判断字符串是否是回文字符串
// 思路：双指针，一个指针i指向初始位置，一个指针j指向末尾位置
//       初始指针i从前向后遍历，末尾指针j从后向前遍历，比较指向的字符是否相等，如果不相等，则必定不是回文字符串
//       循环结束的条件是指针i和j指向同一位置或i>j
func isPalindrome(str string) bool {

	var i, j int

	for i, j = 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			//不相等，则直接返回false
			return false
		}
	}

	return true
}

func main() {

	str := "fsfhgfhg"

	fmt.Printf("str %s is Palindrome: %t \n", str, isPalindrome(str))
}
