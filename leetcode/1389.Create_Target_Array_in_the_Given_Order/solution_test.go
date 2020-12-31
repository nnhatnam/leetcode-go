package _389_Create_Target_Array_in_the_Given_Order

import (
	"github.com/stretchr/testify/assert"

	"testing"
)


func TestSolution(t *testing.T) {
	nums := []int{4,2,4,3,2}
	index := []int{0,0,1,3,1}
	target := createTargetArray(nums, index)
	assert.Equal(t, []int{2,2,4,4,3}, target)
}

