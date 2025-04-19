package main

import "fmt"

// 环形链表

//链表节点定义
type ListNode struct {
	Value byte
	Next  *ListNode
}

//链表
type ByteList struct {
	capacity int
	length   int
	head     *ListNode
}

// 初始化链表
func NewStrList(capacity int) *ByteList {
	return &ByteList{
		capacity: capacity,
		length:   0,
		head:     nil,
	}
}

func (s *ByteList) put(b byte) {
	//维护一个单向链表，链表中加入元素，每次在末尾插入元素
	p := s.head
	q := &ListNode{
		Value: b,
		Next:  nil,
	}
	var pr *ListNode
	if s.head == nil {
		// 链表中无节点
		s.head = q
		s.length++
		return
	}

	// 链表有节点，进行遍历
	for p != nil {
		pr = p
		p = p.Next
	}

	// 如果链表长度大于链表容量，报错返回
	if s.length > s.capacity {
		panic("error")
	}

	// 插入目标节点至链表，其前驱节点为链表尾节点pr
	q = &ListNode{
		Value: b,
		Next:  nil,
	}
	// 插入链表尾部
	pr.Next = q
	// 更新尾节点
	s.length++
}

func hasCycle(head *ListNode) bool {

	// 空链表或只有一个节点，直接返回即可
	if head == nil || head.Next == nil {
		return false
	}

	// 定义快慢指针
	var fast *ListNode
	var slow *ListNode
	fast, slow = head, head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 当快慢指针相遇时，链表肯定有环
		if fast == slow {
			return true
		}
	}

	return false
}

// 环形链表
func main() {
	strList := NewStrList(5)

	strList.put('s')
	strList.put('y')
	strList.put('j')
	strList.put('w')
	strList.put('p')

	var node *ListNode = strList.head
	for ; node.Next != nil; node = node.Next {
		fmt.Printf("%c -> ", node.Value)
	}
	// 链表尾节点指向链表头，构造环形链表
	node.Next = strList.head
	fmt.Printf("List has cycle: %t", hasCycle(strList.head))

}
