package AddTwoNumbers


//Definition for singly-linked list.
type ListNode struct {
     Val int
     Next *ListNode
}

//Use for loop
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	carry , dummy := 0, new(ListNode)
	for tail := dummy; l1 != nil || l2 != nil || carry > 0; tail = tail.Next {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		tail.Next = &ListNode{Val: carry % 10}
		carry /= 10
	}

	return dummy.Next
}


//Use recursion
func addTwoNumbersWithCarry(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 1 {
			return &ListNode{Val: 1}
		}
		return nil
	}

	if l1 != nil {
		carry += l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		carry += l2.Val
		l2 = l2.Next
	}

	node := &ListNode{Val: carry % 10}
	carry /= 10
	node.Next = addTwoNumbersWithCarry(l1, l2, carry)
	return node
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersWithCarry(l1, l2, 0)

	//return dummy.Next
}

//another way to do recursion. it's not a good way to do it since we create a bunch of dummy node and it change l1
func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.Val + l2.Val
	digit , carry := sum % 10 , sum / 10
	result := &ListNode{Val: digit}

	if l1.Next != nil || l2.Next != nil || carry > 0 {
		if l1.Next == nil {
			l1 = &ListNode{Val: 0 }
		} else {
			l1 = l1.Next
		}
		l1.Val += carry

		if l2.Next == nil {
			l2 = &ListNode{Val: 0 }
		} else {
			l2 = l2.Next
		}
		result.Next = addTwoNumbers3(l1, l2)
	}


	return result

	//return dummy.Next
}



//Need to analyze why these are faster (without using dummy node)
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	remindNum := 0
//	pointer1 := l1
//	pointer2 := l2
//	var lhead *ListNode
//	var ltail *ListNode
//	for ;pointer1 != nil || pointer2 != nil ; {
//		var n1, n2 int
//		if pointer1 != nil {
//			n1 = pointer1.Val
//			pointer1 = pointer1.Next
//		}
//		if pointer2 != nil {
//			n2 = pointer2.Val
//			pointer2 = pointer2.Next
//		}
//
//		temp := n1 + n2 + remindNum
//
//		if lhead == nil {
//			lhead = &ListNode{Val: temp % 10}
//			ltail = lhead
//		} else {
//			ltail.Next = &ListNode{Val: temp % 10}
//			ltail = ltail.Next
//		}
//		remindNum = temp / 10
//	}
//
//	if remindNum == 1 {
//		ltail.Next = &ListNode{Val: remindNum}
//		ltail = ltail.Next
//	}
//
//	return lhead
//}

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//
//	resList := &ListNode{}
//	head := resList
//	var i, carry, sum int
//	for l1 != nil || l2 != nil {
//		if l1 == nil{
//			sum = l2.Val + carry
//			l2 = l2.Next
//		}else if l2 == nil{
//			sum = l1.Val + carry
//			l1 = l1.Next
//		} else{
//			sum = l1.Val + l2.Val + carry
//			l1 = l1.Next
//			l2 = l2.Next
//		}
//		carry = sum/10
//		tempN := sum%10
//		if i==0{
//			resList.Val = tempN
//			i++
//			continue
//		}
//		s := &ListNode{
//			Val: tempN,
//		}
//		resList.Next = s
//		resList = resList.Next
//	}
//	if carry != 0 {
//		s := &ListNode{
//			Val: carry,
//		}
//		resList.Next = s
//	}
//	return head
//}
