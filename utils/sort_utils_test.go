package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	elmnt := []int{10, 9, 8, 7, 6, 5}

	BubbleSort(elmnt)

	assert.NotNil(t, elmnt)
	assert.EqualValues(t, 6, len(elmnt))
	assert.EqualValues(t, 5, elmnt[0])
	assert.EqualValues(t, 6, elmnt[1])
	assert.EqualValues(t, 7, elmnt[2])
	assert.EqualValues(t, 8, elmnt[3])
	assert.EqualValues(t, 9, elmnt[4])
	assert.EqualValues(t, 10, elmnt[5])

}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func TestGetElements(t *testing.T) {
	els := getElements(6)
	assert.NotNil(t, els)
	assert.EqualValues(t, 6, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 4, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 2, els[3])
	assert.EqualValues(t, 1, els[4])
	assert.EqualValues(t, 0, els[5])
}

func BenchmarkSort10(b *testing.B) {
	els := getElements(10)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort100(b *testing.B) {
	els := getElements(100)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort10000(b *testing.B) {
	els := getElements(10000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}

func BenchmarkSort100000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(els)
	}
}
