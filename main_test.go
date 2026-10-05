package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	testData := []struct {
		size int
		want int
	}{
		{0, 0},
		{-8, 0},
		{8, 8},
	}

	for _, item := range testData {
		elements := generateRandomElements(item.size)
		assert.Len(t, elements, item.want, "expected: %d, got: %d", item.want, len(elements))
	}
}

func TestMaximum(t *testing.T) {
	testData := []struct {
		data []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{8, 8, 8}, 8},
		{[]int{8000, 50, 300}, 8000},
		{[]int{300, 50, 3, 8000}, 8000},
		{[]int{-100, -50, -20}, -20},
	}

	for _, item := range testData {
		maxNum := maximum(item.data)
		assert.Equal(t, item.want, maxNum, "expected: %d, got: %d", item.want, maxNum)
	}
}

func TestMaxChunks(t *testing.T) {
	testData := []struct {
		data []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{800, 800, 800}, 800},
		{[]int{1, 8, 7, 9000, 1500, 9, 0}, 9000},
		{[]int{1, 8, 7, 9000, 1500, 9, 0, -1}, 9000},
		{[]int{90000, 50, 300}, 90000},
		{[]int{300, 50000, 3, 90000}, 90000},
		{[]int{-100, 0, -20}, 0},
	}

	for _, item := range testData {
		maxChanksVal := maxChunks(item.data)
		assert.Equal(t, item.want, maxChanksVal, "expected: %d, got: %d", item.want, maxChanksVal)
	}
}
