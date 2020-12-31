package _365_How_Many_Numbers_Are_Smaller_Than_the_Current_Number

import (
	"github.com/stretchr/testify/assert"

	"testing"
)


func TestSolution(t *testing.T) {
	t1 := []int{8,1,2,2,3}
	assert.Equal(t, []int{4,0,1,1,3}, smallerNumbersThanCurrent(t1))

	t2 := []int{1,1,1,1,1}
	assert.Equal(t, []int{0,0,0,0,0}, smallerNumbersThanCurrent(t2))
}

