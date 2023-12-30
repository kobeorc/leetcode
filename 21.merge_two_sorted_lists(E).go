package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return m(nil, list1, list2)
}

func m(head, list1, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return head
	}
	if list1 == nil {
		head = add(head, list2.Val)
		list2 = list2.Next
		return m(head, list1, list2)
	}
	if list2 == nil {
		head = add(head, list1.Val)
		list1 = list1.Next
		return m(head, list1, list2)
	}
	if list1.Val > list2.Val {
		head = add(head, list2.Val)
		list2 = list2.Next
		return m(head, list1, list2)
	}
	if list1.Val <= list2.Val {
		head = add(head, list1.Val)
		list1 = list1.Next
		return m(head, list1, list2)
	}
	return head
}

func add(head *ListNode, val int) *ListNode {
	if head == nil {
		head = &ListNode{Val: val}
		return head
	}
	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &ListNode{Val: val}
	return head
}
