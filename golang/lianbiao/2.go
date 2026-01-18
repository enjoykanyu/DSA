package lianbiao

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var value *ListNode
	jin := 0 //进位
	for l1 != nil && l2 != nil {
		val := (l1.Val + l2.Val) + jin
		l1 = l1.Next
		l2 = l2.Next

		jin = val / 10
		if head == nil {
			head = &ListNode{Val: val % 10}
			value = head
		} else {
			value.Next = &ListNode{Val: val % 10}
			value = value.Next
		}
	}
	for l1 != nil {
		val := (l1.Val) + jin
		l1 = l1.Next
		jin = val / 10
		if head == nil {
			head = &ListNode{Val: val % 10}
			value = head
		} else {
			value.Next = &ListNode{Val: val % 10}
			value = value.Next
		}
	}
	for l2 != nil {
		val := (l2.Val) + jin
		l2 = l2.Next
		jin = val / 10
		if head == nil {
			head = &ListNode{Val: val % 10}
			value = head
		} else {
			value.Next = &ListNode{Val: val % 10}
			value = value.Next
		}
	}
	if jin > 0 {
		value.Next = &ListNode{Val: jin}
	}
	return head
}
