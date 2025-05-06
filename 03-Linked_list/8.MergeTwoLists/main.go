package main

import "fmt"

// 链表中间节点

//链表节点定义
type ListNode struct {
	Value int
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

func (s *ByteList) put(b int) {
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

// 合并有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	// 引入哑节点
	dummy := &ListNode{}
	// 当前指针
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Value <= list2.Value {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 == nil {
		cur.Next = list2
	} else {
		cur.Next = list1
	}

	return dummy.Next
}

// 链表中间节点
func main() {
	strList1 := NewStrList(6)

	strList1.put('1')
	strList1.put('3')
	strList1.put('4')
	strList1.put('4')
	strList1.put('8')
	strList1.put('9')

	var node1 *ListNode = strList1.head
	for ; node1 != nil; node1 = node1.Next {
		fmt.Printf("%c -> ", node1.Value)
	}
	fmt.Println("")

	strList2 := NewStrList(6)

	strList2.put('1')
	strList2.put('2')
	strList2.put('4')

	var node2 *ListNode = strList2.head
	for ; node2 != nil; node2 = node2.Next {
		fmt.Printf("%c -> ", node2.Value)
	}

	fmt.Println("")
	var node *ListNode = mergeTwoLists(strList1.head, strList2.head)
	for ; node != nil; node = node.Next {
		fmt.Printf("%c -> ", node.Value)
	}
	fmt.Println("")

}
