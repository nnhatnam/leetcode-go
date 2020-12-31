package _512_Number_of_Good_Pairs

import (
	"github.com/stretchr/testify/assert"

	"testing"
)


func TestSolution(t *testing.T) {
	t1 := []int{1,1,1,1}
	assert.Equal(t, numIdenticalPairs(t1), 6)
}

