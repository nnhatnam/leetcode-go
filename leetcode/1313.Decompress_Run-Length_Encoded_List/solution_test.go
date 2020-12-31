package _313_Decompress_Run_Length_Encoded_List

import (
	"github.com/stretchr/testify/assert"

	"testing"
)


func TestSolution(t *testing.T) {
	t1 := []int{1,2,3,4}
	assert.Equal(t, decompressRLElist(t1), []int{2,4,4,4})
}


