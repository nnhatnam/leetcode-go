package _1_Merge_Two_Sorted_Lists


type ListNode struct {
     Val int
 	Next *ListNode
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	head := &ListNode{}
	current := head

	for ; l1 != nil && l2 != nil ; {

		if  l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
			current = current.Next
		} else {
			current.Next = l2
			l2 = l2.Next
			current = current.Next
		}

	}
	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}

	return head.Next
}
