package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	l1.Val = 1
	l1.Next = new(ListNode)
	l1.Next.Val = 2
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 4

	l2 := new(ListNode)
	l2.Val = 1
	l2.Next = new(ListNode)
	l2.Next.Val = 3
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4
	d := mergeTwoLists(l1, l2)
	fmt.Printf("ddddddddddddddddd=%v\n", *d)
	fmt.Printf("ddddddddddddddddd=%v\n", *d.Next)
	fmt.Printf("ddddddddddddddddd=%v\n", *d.Next.Next)
	fmt.Printf("ddddddddddddddddd=%v\n", *d.Next.Next.Next)
	fmt.Printf("ddddddddddddddddd=%v\n", *d.Next.Next.Next.Next)
	fmt.Printf("ddddddddddddddddd=%v\n", *d.Next.Next.Next.Next.Next)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	merge(newHead, l1, l2)
	return newHead.Next
}
func merge(head *ListNode, l1 *ListNode, l2 *ListNode) {
	if l1 == nil && l2 == nil {
		return
	}
	if l1 == nil {
		head.Next = l2
		return
	}
	if l2 == nil {
		head.Next = l1
		return
	}
	if l1.Val <= l2.Val {
		head.Next = l1
		merge(head.Next, l1.Next, l2)
	} else {
		head.Next = l2
		merge(head.Next, l1, l2.Next)
	}
	return
}
