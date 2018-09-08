package searching

import "testing"

var searchTests = []struct {
	input    []int
	find     int
	expected int
}{
	//{[]int{}, 0, -1},                                                                                                                  // return -1 if array is empty
	//{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 4, -1},                                                                                 // return false if not in array of odd numbers
	//{[]int{2, 4, 6, 8, 10, 12, 14, 16, 18}, 5, -1},                                                                                    // return false if not in array of even numbers
	{[]int{5, 4, 3, 2, 1}, 5, 0},                                                                                                      // find first entry in array
	//{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 8},                                                                                          // find last entry in array
	//{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 4},                                                                                          // find entry in middle of array
	//{[]int{38, 98, 2, 47, 38, 647, 28, 3, 3, 74, 8, 32, 86, 3, 5, 89, 7, 3, 78, 49, 8, 0, 5, 43, 21, 7, 654, 6, 16, 8, 7, 34}, 89, 15}, // fin entry anywhere in array
	//{[]int{38, 98, 2, 47, 38, 647, 28, 3, 3, 74}, 1, -1},                                                                              // return -1 if not in array
}

func TestLinearSearch(t *testing.T) {
	for _, tt := range searchTests {
		actual := LinearSearch(tt.input, tt.find)
		if actual != tt.expected {
			t.Errorf("LinearSearch: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("LinearSearch: expected %d, actual %d", tt.expected, actual)
	}
}

func TestBinarySearch(t *testing.T) {
	for _, tt := range searchTests {
		actual := BinarySearch(tt.input, tt.find)
		if actual != tt.expected {
			t.Errorf("BinarySearch: expected %d, actual %d", tt.expected, actual)
		}
		//t.Logf("BinarySearch: expected %d, actual %d", tt.expected, actual)
	}
}

func TestInterpolationSearch(t *testing.T) {
	for _, tt := range searchTests {
		actual := InterpolationSearch(tt.input, tt.find)
		if actual != tt.expected {
			t.Errorf("InterpolationSearch: expected %d, actual %d", tt.expected, actual)
		}
		t.Logf("InterpolationSearch: expected %d, actual %d", tt.expected, actual)
	}
}
