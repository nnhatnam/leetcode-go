package _2_Remove_Duplicates_from_Sorted_List_II


type ListNode struct {
     Val int
     Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

	sentinel := &ListNode{}
	sentinel.Next = head
	previous := sentinel

	for ; head != nil; {
		if head.Next != nil && head.Val == head.Next.Val {

			for ; head.Next != nil && head.Val == head.Next.Val ; {
				head = head.Next
			}

			previous.Next = head.Next
		} else {
			previous = previous.Next
		}
		head = head.Next
	}

	return sentinel.Next
}