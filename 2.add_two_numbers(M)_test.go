package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val:  9,
		Next: nil,
	}
	l2 := &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	}
	lr := &ListNode{
		Val: 8,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  0,
				Next: &ListNode{Val: 1},
			},
		},
	}

	ltc := addTwoNumbers(l1, l2)
	if !compare(ltc, lr) {
		t.Logf("%+v", printList(ltc))
		t.Fail()
		return
	}
	t.Log("Success")
}
