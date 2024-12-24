package mergesort

/*
Perform acending order merge sort on a slice of intergers
*/
func PerformMergeSort(numbers []int) []int {
	n := len(numbers)
	sorted_list := make([]int, 0, n)

	sorted_list = append(sorted_list, merge_arr(PerformMergeSort(numbers[:n/2]), PerformMergeSort(numbers[n/2:n]))...)

	return sorted_list
}

/*
Take the contents of two equal sized lists and merge them in ascending order, two finger style
*/
func merge_arr(a, b []int) []int {
	if len(a) == 1 && len(b) == 1 {
		if a[0] > b[0] {
			return []int{b[0], a[0]}
		} else {

			return []int{a[0], b[0]}
		}
	}

	sorted_array := make([]int, 0, len(a)+len(b))
	p1 := 0
	p2 := 0

	for i := 0; i < len(a)+len(b); i++ {
		if a[p1] > b[p2] {
			sorted_array = append(sorted_array, b[p2])
			p2 += 1
		} else {
			sorted_array = append(sorted_array, a[p1])
			p1 += 1
		}
	}

	return sorted_array
}
