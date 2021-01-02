package _282_Group_the_People_Given_the_Group_Size_They_Belong_To

import (
	"github.com/stretchr/testify/assert"

	"testing"
)


func TestSolution(t *testing.T) {
	t1 := []int{3,3,3,3,3,1,3}
	assert.Equal(t, [][]int{[]int{0, 1, 2}, []int{5}, []int{3, 4, 6}} , groupThePeople(t1))
}

