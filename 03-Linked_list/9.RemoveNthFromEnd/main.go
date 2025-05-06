package main

import "fmt"

// 链表中间节点

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var cur *ListNode = head
	var listLen int = 0
	// 获取链表长度
	for cur != nil {
		listLen++
		cur = cur.Next
	}

	// 待删除节点是链表头节点时，先保存头节点的后继节点，删除头节点，返回
	if n == listLen {
		var newHead *ListNode = head.Next
		head.Next = nil
		return newHead
	}
	// 当cur走到 (listLen - n)步时，为待删除节点的前躯节点
	var i int = 1
	for cur = head; i < listLen-n; cur = cur.Next {
		i++
	}

	// 特殊情况，待删除节点是尾节点
	if cur.Next.Next == nil {
		cur.Next = nil
		return head
	}
	// 待删除节点的前驱节点指向待删除节点后继节点
	cur.Next = cur.Next.Next
	return head
}

func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	// 快慢指针
	var fast, slow *ListNode = head, head
	for i := 1; i <= n; fast = fast.Next {
		i++
	}
	if fast == nil {
		var newHead *ListNode = head.Next
		head.Next = nil
		return newHead
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 特殊情况，待删除节点是尾节点
	if slow.Next.Next == nil {
		slow.Next = nil
		return head
	}

	slow.Next = slow.Next.Next
	return head
}

// 移除链表倒数第n个节点
func main() {
	strList := NewStrList(6)

	strList.put('s')
	strList.put('y')
	strList.put('j')
	strList.put('w')
	strList.put('p')
	strList.put('q')

	var node *ListNode = strList.head
	for ; node != nil; node = node.Next {
		fmt.Printf("%c -> ", node.Value)
	}
	fmt.Println("")
	var newHead *ListNode = removeNthFromEnd(strList.head, 6)

	for ; newHead != nil; newHead = newHead.Next {
		fmt.Printf("%c -> ", newHead.Value)
	}
	fmt.Println("")
	fmt.Println("Call removeNthFromEndV2")
	strList2 := NewStrList(6)

	strList2.put('s')
	strList2.put('y')
	strList2.put('j')
	strList2.put('w')
	strList2.put('p')
	strList2.put('q')

	newHead = removeNthFromEndV2(strList2.head, 6)

	for ; newHead != nil; newHead = newHead.Next {
		fmt.Printf("%c -> ", newHead.Value)
	}
}
