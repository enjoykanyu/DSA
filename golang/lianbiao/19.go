package lianbiao

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//两次for循环
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count := 0
	pre := &ListNode{Next: head, Val: 1}
	cur := &ListNode{Next: head, Val: 1}
	temp := head
	for temp != nil {
		count++
		temp = temp.Next
	}
	value := count - n
	t := 0
	if n == count { //长度等于删除索引 剔除第0个数据
		return head.Next
	}
	for t < value {
		cur = cur.Next
		t++
	}
	cur.Next = cur.Next.Next
	return pre.Next
}
