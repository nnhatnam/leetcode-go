package _32_Flipping_an_Image

import (
	"github.com/stretchr/testify/assert"

	"testing"
)


func TestSolution(t *testing.T) {
	t1 := [][]int{{1,1,0},{1,0,1},{0,0,0}}
	assert.Equal(t, [][]int{{1,0,0},{0,1,0},{1,1,1}} , flipAndInvertImage(t1))

	t2 := [][]int{{1}}
	assert.Equal(t, [][]int{{0}} , flipAndInvertImage(t2))
}
