package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		sum := n1 + n2 + carry
		if sum > 9 {
			sum = sum - 10
			carry = 1
		} else {
			carry = 0
		}
		curr.Next = &ListNode{Val: sum, Next: nil}
		curr = curr.Next
	}
	return head.Next
}
