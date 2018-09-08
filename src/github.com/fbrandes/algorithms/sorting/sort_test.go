package sorting

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// Test values
var sortTests = []struct {
	values   []int // input
	expected []int // expected result
}{
	{[]int{3, 6, 5, 79, 8, 8, 498, 47, 5}, []int{3, 5, 5, 6, 8, 8, 47, 79, 498}},                     // test with repeating values
	{[]int{3, -6, 5, -79, -8, -68, 498, 47, -5}, []int{-79, -68, -8, -6, -5, 3, 5, 47, 498}},         // test with positive and negative values
	{[]int{-3, -6, -5, -79, -8, -68, -498, -47, -5}, []int{-498, -79, -68, -47, -8, -6, -5, -5, -3}}, // test with only negative values
	{[]int{3, 6, 5, 79, 8, 8, 498, 47, 5}, []int{3, 5, 5, 6, 8, 8, 47, 79, 498}},                     // test with only positive numbers
	{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},                             // test with reverse ordered input
	{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},                                                           // test with sorted input
	{[]int{4}, []int{4}},                                                                             // test with one element only
	{[]int{}, []int{}},                                                                               // test with empty input
}

// Test Cases
func TestBubbleSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := BubbleSort(tt.values)
		assert.Equal(t, actual, tt.expected, "BubbleSort: expected %d, actual %d")
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := MergeSort(tt.values)
		assert.Equal(t, actual, tt.expected, "MergeSort: expected %d, actual %d")
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := QuickSort(tt.values)
		assert.Equal(t, actual, tt.expected, "QuickSort: expected %d, actual %d")
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := InsertionSort(tt.values)
		assert.Equal(t, actual, tt.expected, "InsertionSort: expected %d, actual %d")
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := SelectionSort(tt.values)
		assert.Equal(t, actual, tt.expected, "SelectionSort: expected %d, actual %d")
	}
}

func TestHeapSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := HeapSort(tt.values)
		assert.Equal(t, actual, tt.expected, "HeapSort: expected %d, actual %d")
	}
}

func TestRadixSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := RadixSort(tt.values)
		assert.Equal(t, actual, tt.expected, "RadixSort: expected %d, actual %d")
	}
}

func TestShellSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := ShellSort(tt.values)
		assert.Equal(t, actual, tt.expected, "ShellSort: expected %d, actual %d")
	}
}

func TestCountingSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := CountingSort(tt.values)
		assert.Equal(t, actual, tt.expected, "CountingSort: expected %d, actual %d")
	}
}

func TestBucketSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := BucketSort(tt.values)
		assert.Equal(t, actual, tt.expected, "BucketSort: expected %d, actual %d")
	}
}

func TestCombSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := CombSort(tt.values)
		assert.Equal(t, actual, tt.expected, "CombSort: expected %d, actual %d")
	}
}
