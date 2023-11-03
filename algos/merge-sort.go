package algos

//
//has a time complexity of O(n log n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	//split the array in half
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	//merge the two sorted halves
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	//append any remaining elements(if any)
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
