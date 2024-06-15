package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {

	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8}

	BubbleSort(elements)

	assert.NotNil(t, elements)
	assert.EqualValues(t, 9, len(elements))
	assert.EqualValues(t, 1, elements[0])
	assert.EqualValues(t, 2, elements[1])
	assert.EqualValues(t, 3, elements[2])
	assert.EqualValues(t, 4, elements[3])
	assert.EqualValues(t, 5, elements[4])
	assert.EqualValues(t, 6, elements[5])
	assert.EqualValues(t, 7, elements[6])
	assert.EqualValues(t, 8, elements[7])
	assert.EqualValues(t, 9, elements[8])
}

func TestBubbleSortNilSlice(t *testing.T) {

	BubbleSort(nil)

}

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8}

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}

}

func BenchmarkBubbleSortWorstCase(b *testing.B) {
	elements := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}

}
