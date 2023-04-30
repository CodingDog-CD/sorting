// package quick_sort provides the implementation of quick sort algorithm
package quick_sort

// Quick sort algorithm used to sort a series of numbers.
// Time: O(nlogn), Space: O(log(n)). Space complexity is relevant to the height of the recursion tree.
func Sort(input []int, left_idx int, right_idx int) {
	if left_idx < right_idx {
		partitionIndex := partition(input, left_idx, right_idx)
		if partitionIndex > (left_idx+right_idx)/2 {
			Sort(input, partitionIndex+1, right_idx)
			Sort(input, left_idx, partitionIndex-1)
		} else {
			Sort(input, left_idx, partitionIndex-1)
			Sort(input, partitionIndex+1, right_idx)
		}
	}
}

// Partition function is used to put the pivot into the correct position and return it's index in the modified array
func partition(input []int, left_idx int, right_idx int) int {
	pivot := input[left_idx]
	for left_idx < right_idx {
		for input[right_idx] >= pivot && left_idx < right_idx {
			right_idx--
		}
		if left_idx < right_idx {
			input[left_idx] = input[right_idx]
			left_idx++
		}
		for input[left_idx] <= pivot && left_idx < right_idx {
			left_idx++
		}
		if left_idx < right_idx {
			input[right_idx] = input[left_idx]
			right_idx--
		}
	}
	input[left_idx] = pivot
	return left_idx
}
