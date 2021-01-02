package _7_Remove_Element

import (
	"github.com/stretchr/testify/assert"

	"testing"
)


func TestSolution(t *testing.T) {
	nums := []int{3,2,2,3}
	val := 3
	assert.Equal(t, 2, removeElement(nums, val))
}


