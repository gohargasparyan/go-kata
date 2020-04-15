package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKarateChop(t *testing.T, chop func(int, []int) int)  {
	assert.Equal(t, -1, chop(3, []int{}))

	assert.Equal(t, -1, chop(3, []int{1}))
	assert.Equal(t, 0, chop(1, []int{1}))

	assert.Equal(t, 0, chop(1, []int{1, 3, 5}))
	assert.Equal(t, 1, chop(3, []int{1, 3, 5}))
	assert.Equal(t, 2, chop(5, []int{1, 3, 5}))
	assert.Equal(t, -1, chop(0, []int{1, 3, 5}))
	assert.Equal(t, -1, chop(4, []int{1, 3, 5}))
	assert.Equal(t, -1, chop(6, []int{1, 3, 5}))

	assert.Equal(t, 6, chop(63, []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 101}))
	assert.Equal(t, -1, chop(61, []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 101}))
	assert.Equal(t, -1, chop(200, []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 101}))
	assert.Equal(t, -1, chop(-5, []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 101}))
	assert.Equal(t, 0, chop(1, []int{1, 2, 9, 20, 31, 45, 63, 70, 100, 101}))
}
