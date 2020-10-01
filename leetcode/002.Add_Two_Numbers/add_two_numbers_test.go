package AddTwoNumbers

import (
"testing"
)


func arrToList(arr []int) *ListNode {
	var head, tail *ListNode
	for _, v := range arr {
		if head == nil {
			head = &ListNode{Val: v}
			tail = head
		} else {
			tail.Next = &ListNode{Val: v}
			tail = tail.Next
		}
	}
	return head
}

func listToArr(l *ListNode) []int {
	result := []int{}
	curr := l
	for ; curr != nil ; curr = curr.Next{
		result = append(result, curr.Val)
	}
	return result

}

type testCase struct {
	inputs
	outputs
}

type inputs struct {
	param1 []int
	param2 []int
}

type outputs struct {
	result []int
}

//Input: [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1] , [5,6,4]
//Expected: [6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
func TestEdgeCases(t *testing.T) {
	var testcases = []testCase{
		{
			inputs{param1: []int{0} , param2: []int{0}},
			outputs{result : []int{0}},
		},

		{
			inputs{param1: []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1} , param2: []int{5,6,4}},
			outputs{result : []int{6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}},
		},
	}

	for _ , tcase := range testcases {
		result := addTwoNumbers2(arrToList(tcase.param1), arrToList(tcase.param2))
		index := 0
		t.Log("Output   : ", listToArr(result))
		t.Log("Expected : ", tcase.outputs.result)
		for ; result != nil ;result = result.Next {

			if tcase.outputs.result[index] != result.Val {
				t.Errorf("Expected : %v, got %v", tcase.outputs.result, listToArr(result))
				break
			}
			index += 1
		}
	}

}

