package lianbiao

// * Definition for singly-linked list.
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	listA := headA
	listB := headB
	for listA != nil {
		listB = headB
		for listB != nil {
			if listA != listB {
				listB = listB.Next
			} else {
				return listB
			}
		}
		listA = listA.Next
	}
	return nil
}
