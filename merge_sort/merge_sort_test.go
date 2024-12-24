package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	input := []int{6, 5, 3, 1, 8, 7, 2, 4}
	expected_output := []int{1, 2, 3, 4, 5, 6, 7, 8}

	actual_output := PerformMergeSort(input)

	if !reflect.DeepEqual(expected_output, actual_output) {
		t.Fatalf("Actual output of merge sort: %v\n, not the same as expected %v", actual_output, expected_output)
	}
}
