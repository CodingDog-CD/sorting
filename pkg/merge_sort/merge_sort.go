// package merge_sort provide the implementation of merge sort algorithm
package merge_sort

// merge sort.
// Time: O(nlog(n)), Space: O(log(n))
func Sort(input []int) []int {
	length := len(input)
	middle := (len(input) / 2)
	if middle > 0 {
		return merge(Sort(input[0:middle]), Sort(input[middle:length]))
	} else {
		return input
	}

}

func merge(leftSubarray, rightSubarray []int) []int {
	var mergedArray []int
	if len(leftSubarray) == 1 && len(rightSubarray) == 1 {
		if leftSubarray[0] < rightSubarray[0] {
			mergedArray = append(mergedArray, leftSubarray[0], rightSubarray[0])
		} else {
			mergedArray = append(mergedArray, rightSubarray[0], leftSubarray[0])
		}
	} else {
		idxLeft := 0
		idxRight := 0
		for idxLeft < len(leftSubarray) && idxRight < len(rightSubarray) {
			if leftSubarray[idxLeft] < rightSubarray[idxRight] {
				mergedArray = append(mergedArray, leftSubarray[idxLeft])
				idxLeft++
			} else {
				mergedArray = append(mergedArray, rightSubarray[idxRight])
				idxRight++
			}
		}
		if idxLeft >= len(leftSubarray) {
			mergedArray = append(mergedArray, rightSubarray[idxRight:]...)
		} else {
			mergedArray = append(mergedArray, leftSubarray[idxLeft:]...)
		}

	}
	return mergedArray
}
