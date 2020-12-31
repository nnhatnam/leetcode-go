package _588_Sum_of_All_Odd_Length_Subarrays

import (
	"github.com/stretchr/testify/assert"

	"testing"
)


func TestSolution(t *testing.T) {
	t1 := []int{1,4,2,5,3}
	assert.Equal(t, 58, sumOddLengthSubarrays(t1) )
}


