package sorting

import (
	"testing"
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
		if !checkAccuracy(actual, tt.expected) {
			t.Errorf("BubbleSort: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("BubbleSort: expected %d, actual %d", tt.expected, actual)
	}
}

func TestMergeSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := MergeSort(tt.values)
		if !checkAccuracy(actual, tt.expected) {
			t.Errorf("MergeSort: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("MergeSort: expected %d, actual %d", tt.expected, actual)
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := QuickSort(tt.values)
		if !checkAccuracy(actual, tt.expected) {
			t.Errorf("QuickSort: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("QuickSort: expected %d, actual %d", tt.expected, actual)
	}
}

func TestInsertionSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := InsertionSort(tt.values)
		if !checkAccuracy(actual, tt.expected) {
			t.Errorf("InsertionSort: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("InsertionSort: expected %d, actual %d", tt.expected, actual)
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := SelectionSort(tt.values)
		if !checkAccuracy(actual, tt.expected) {
			t.Errorf("SelectionSort: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("SelectionSort: expected %d, actual %d", tt.expected, actual)
	}
}

func TestHeapSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := HeapSort(tt.values)
		if !checkAccuracy(actual, tt.expected) {
			t.Errorf("HeapSort: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("HeapSort: expected %d, actual %d", tt.expected, actual)
	}
}

func TestRadixSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := RadixSort(tt.values)
		if !checkAccuracy(actual, tt.expected) {
			t.Errorf("RadixSort: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("RadixSort: expected %d, actual %d", tt.expected, actual)
	}
}

func TestShellSort(t *testing.T) {
	for _, tt := range sortTests {
		actual := ShellSort(tt.values)
		if !checkAccuracy(actual, tt.expected) {
			t.Errorf("ShellSort: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("ShellSort: expected %d, actual %d", tt.expected, actual)
	}
}

// Assertion Helpers
func checkAccuracy(source []int, result []int) bool {
	return (len(source) == len(result)) && allValuesPresent(source, result) && isOrderd(result)
}

func isOrderd(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if !(a[i] <= a[i+1]) {
			return false
		}
	}
	return true
}
func allValuesPresent(source []int, result []int) bool {
	for i := 0; i < len(source); i++ {
		if !(contains(result, source[i])) {
			return false
		}
	}
	return true
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
