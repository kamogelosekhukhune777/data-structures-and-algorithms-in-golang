package algos

//SelectionSort repeatedly selects the minimum(or maximum)
// element from the unsorted part of the array and puts it
// at the beginning(or end).
//
//has a time complexity of O(n^2).
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		//find the minimum element in the unsorted part of the array
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//swap the found minimum element with the current element
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
