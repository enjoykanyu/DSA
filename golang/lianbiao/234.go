package lianbiao

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	list := []int{}
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	// 正确的双指针双向遍历
	for t, l := 0, len(list)-1; t <= l; {
		if list[t] != list[l] {
			return false
		}
		t++
		l--
	}
	return true
}
