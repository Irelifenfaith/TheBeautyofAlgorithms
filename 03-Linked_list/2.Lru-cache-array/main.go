package main

import "fmt"

// Get方法：查找并更新LRU
func Get(cache *[3]int, value int) bool {
	// 思路：维护一个有序数组，越靠近数组尾部的是越早访问的。
	// 当一个新数据访问时，从数组头顺序遍历数组
	// 1.如果数据已经被缓存到数组中，则将该节点插入到数组头部；
	// 2.如果数据在数组中未找到（由于数组容量是固定的，因此只有一种情况）：
	//   a.删除最后一个元素，将数据插入到头部
	//总时间复杂度:
	// 最坏情况：无论是否找到目标值，都需要遍历数组并移动元素：
	// 查找 + 移动：O(n) + O(n) = O(n)。
	// 平均情况：同样为 O(n)，因为每次操作都需要遍历数组的一部分。
	// 因此时间复杂度是O(n)
	var end int
	p := *cache
	length := len(*cache)
	found := false
	for i := 0; i < length; i++ {
		if p[i] == value {
			end = i
			found = true
			break
		}
	}

	if found {
		// 查到了p[i]即查到的元素
		// 查到了，如果是头元素，无需移动直接返回(end为0，说明上面循环一次就找到了)
		if end == 0 {
			return found
		}
		//非头元素
		//挪动位置（[2,3,1,5]，查到了3，则前一元素是2，后面是1，需变为[3,2,1,5]）
		//将0到i的元素进行后移
		for k := end; k > 0; k-- {
			(*cache)[k] = (*cache)[k-1]
		}
		// 将值插入到数组首元素
		(*cache)[0] = value
		return found
	} else {
		// 没查到：将元素插入到数组初始位置；将数组最后一个元素淘汰
		for k := len(*cache) - 1; k > 0; k-- {
			(*cache)[k] = (*cache)[k-1]
		}
		(*cache)[0] = value
		return false
	}

}

// LRU Cache
func main() {
	// LRU Cache初始化
	var cache [3]int
	fmt.Println(Get(&cache, 1))
	fmt.Printf("验证头节点找到：")
	fmt.Println(Get(&cache, 1))
	Get(&cache, 2)
	// 验证数组结构
	fmt.Println("Get(2)后链表顺序：", cache)
	Get(&cache, 3)
	// 验证数组结构
	fmt.Println("Get(3)后链表顺序：", cache)

	Get(&cache, 4)
	// 验证数组结构
	fmt.Println("Get(4)后链表顺序：", cache)

	Get(&cache, 2)
	// 验证数组结构
	fmt.Println("Get(2)后链表顺序：", cache)

	Get(&cache, 5)
	// 验证数组结构
	fmt.Println("Get(5)后链表顺序：", cache)
	//输出结果如下：
	// false
	// 验证头节点找到：true
	// Get(2)后链表顺序： [2 1 0]
	// Get(3)后链表顺序： [3 2 1]
	// Get(4)后链表顺序： [4 3 2]
	// Get(2)后链表顺序： [2 4 3]
	// Get(5)后链表顺序： [5 2 4]

}
