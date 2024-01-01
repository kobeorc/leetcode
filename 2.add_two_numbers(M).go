package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(nil, l1, l2, 0)
}

func addTwo(head, l1, l2 *ListNode, plus int) *ListNode {
	if l1 == nil && l2 == nil {
		if plus >= 1 {
			tmp := head
			for tmp.Next != nil {
				tmp = tmp.Next
			}

			tmp.Next = &ListNode{
				Val:  plus,
				Next: nil,
			}

		}
		return head
	}
	if l1 == nil {
		v := l2.Val + plus
		head = addAddtwonumbers(head, 0, v%10)
		l2 = l2.Next
		return addTwo(head, l1, l2, v/10)
	}

	if l2 == nil {

		v := l1.Val + plus
		head = addAddtwonumbers(head, v%10, 0)
		l1 = l1.Next
		return addTwo(head, l1, l2, v/10)
	}

	head = addAddtwonumbers(head, l1.Val, l2.Val+plus)
	t := (l1.Val + l2.Val + plus) / 10
	l1 = l1.Next
	l2 = l2.Next

	return addTwo(head, l1, l2, t)
}

func addAddtwonumbers(head *ListNode, l1, l2 int) *ListNode {
	if head == nil {
		return &ListNode{
			Val:  (l1 + l2) % 10,
			Next: nil,
		}

	}

	tmp := head
	for tmp.Next != nil {
		tmp = tmp.Next
	}

	tmp.Next = &ListNode{
		Val:  (l1 + l2) % 10,
		Next: nil,
	}

	return head
}
