package main

import "fmt"

//链表节点定义
type ListNode struct {
	value byte
	next  *ListNode
	prev  *ListNode
}

//双向链表
type ByteList struct {
	capacity int
	length   int
	head     *ListNode
	tail     *ListNode
}

// 初始化链表
func NewStrList(capacity int) *ByteList {
	return &ByteList{
		capacity: capacity,
		length:   0,
		head:     nil,
		tail:     nil,
	}
}

func (s *ByteList) put(b byte) {
	//维护一个双向链表，链表中加入元素，每次在末尾插入元素
	p := s.head
	q := &ListNode{
		value: b,
		next:  nil,
	}
	var pr *ListNode
	if s.head == nil {
		// 链表中无节点
		s.head = q
		s.tail = q
		s.length++
		return
	}

	// 链表有节点，进行遍历
	for p != nil {
		pr = p
		p = p.next
	}

	// 如果链表长度大于链表容量，报错返回
	if s.length > s.capacity {
		panic("error")
	}

	// 插入目标节点至链表，其前驱节点为链表尾节点pr
	q = &ListNode{
		value: b,
		next:  nil,
		prev:  pr,
	}
	// 插入链表尾部
	s.tail.next = q
	// 更新尾节点
	s.tail = q
	s.length++
}

func revertList(head *ListNode) *ListNode {

	// 当前第一个节点
	var pr *ListNode
	var next *ListNode
	// 当前节点，初始为head
	cur := head

	// 空链表或只有一个节点，直接返回即可
	if head == nil || head.next == nil {
		return nil
	}

	for cur != nil {
		// 保存当前节点的后继节点
		next = cur.next
		// 将当前节点的后继节点指向前驱节点
		cur.next = pr
		// 移动pr到当前节点
		pr = cur
		// 移动cur到后继节点
		cur = next
		// fmt.Printf("cur.value: %c, pr.value: %c\n", cur.value, pr.value)
	}
	head = pr

	return head
}

// 判断链表是否是回文链表（单链表）
// 思路：快慢指针，反转链表后半部分,然后进行比较
func isPalindromeSingleList(head *ListNode) bool {
	slow, fast := head, head

	if head == nil || head.next == nil {
		return true
	}

	// 快慢指针，slow走一步，fast走两步，fast走到链表尾部时，slow为链表中间位置
	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next
	}

	fmt.Println("go ")
	// 反转链表后半部分
	newHead := revertList(slow)
	for head != nil || newHead != nil {
		if head.value != newHead.value {
			return false
		}
		head = head.next
		newHead = newHead.next
	}

	return true
}

func getTail(head *ListNode) *ListNode {
	current := head
	for current.next != nil {
		current = current.next
	}
	return current
}

// 判断链表是否是回文链表（双向链表）
// 思路：双指针，一个指针i指向初始位置，一个指针j指向末尾位置
//       初始指针i从前向后遍历，末尾指针j从后向前遍历，比较指向的值是否相等，如果不相等，则必定不是回文链表
//       循环结束的条件是指针i和j指向同一位置或i>j
func isPalindromeDoubleList(head *ListNode) bool {

	var tail *ListNode = getTail(head)
	left, right := head, tail

	for left != nil && right != nil && left != right && left.prev != right {
		if head.value != tail.value {
			return false
		}
		fmt.Printf("head.value is %c, tail.value is %c\n", head.value, tail.value)
		left = left.next
		right = right.prev
	}

	return true
}

// 判断链表是否是回文链表
func main() {
	strList := NewStrList(7)

	strList.put('s')
	strList.put('f')
	strList.put('g')
	strList.put('f')
	strList.put('s')

	for node := strList.head; node != nil; node = node.next {
		fmt.Printf("%c <-> ", node.value)
	}
	fmt.Println("")
	fmt.Printf("strList is PalindromeDoubleList: %t \n", isPalindromeDoubleList(strList.head))
	fmt.Printf("strList is PalindromeSingleList: %t \n", isPalindromeSingleList(strList.head))
}
