package Check_Array_Formation_Through_Concatenation

import (

	"testing"
)


func TestSolution(t *testing.T) {
	arr := []int{37,69,3,74,46}
	pieces := [][]int{{37,69,3,74,46}}
	canFormArray(arr, pieces)

}


