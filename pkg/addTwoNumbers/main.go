package addTwoNumber

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	over := false
	var ans *ListNode
	var n1, n2 = l1, l2
	for i := 0; n1 != nil || n2 != nil; i++ {
		digit := 0
		if n1 != nil {
			digit += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			digit += n2.Val
			n2 = n2.Next
		}
		if over {
			digit += 1
		}
		if digit > 9 {
			digit -= 10
			over = true
		} else {
			over = false
		}
		ans = &ListNode{Val: digit, Next: ans}
	}
	if over {
		ans = &ListNode{Val: 1, Next: ans}
	}
	// 反転して返す
	var prev *ListNode
	for ans != nil {
		next := ans.Next
		ans.Next = prev
		prev = ans
		ans = next
	}
	return prev
}

func Ans() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	ans := addTwoNumbers(l1, l2)
	for ans != nil {
		fmt.Println(ans.Val)
		ans = ans.Next
	}
}
