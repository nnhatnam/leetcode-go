package _539_Kth_Missing_Positive_Number

import (
	"github.com/stretchr/testify/assert"

	"testing"
)


func TestSolution(t *testing.T) {
	arr := []int{2,3,4,7,11}
	k := 5
	assert.Equal(t, 9 , findKthPositive(arr, k))

	arr = []int{1,2,3,4}
	k = 2
	assert.Equal(t, 6 , findKthPositive(arr, k))
}


