package main

import (
	"fmt"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	l2 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	ltc := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	lr := mergeTwoLists(&l1, &l2)

	if !compare(ltc, lr) {
		t.Logf("%+v", printList(lr))
		t.Fail()
		return
	}
	t.Log("Success")
}

func compare(l1, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if (l1 == nil && l2 != nil) || (l1 != nil && l2 == nil) {
		return false
	}
	if l1.Val != l2.Val {
		return false
	}

	return compare(l1.Next, l2.Next)
}

var calcList string

func printList(l *ListNode) string {
	calcList += fmt.Sprint(l.Val, " -> ")
	if l.Next == nil {
		return calcList
	}
	return printList(l.Next)
}
