package algos

//QuickSort
//
//has a time complexity of O(n log n)
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivotIndex := len(arr) / 2
	pivot := arr[pivotIndex]

	var less, equal, greather []int
	for _, num := range arr {
		if num < pivot {
			less = append(less, num)
		} else if num == pivot {
			equal = append(equal, num)
		} else {
			greather = append(greather, num)
		}
	}

	//recursively sort the partitions
	less = QuickSort(less)
	greather = QuickSort(greather)
	//combine the sortes partitions
	sortedArr := append(less, equal...)
	sortedArr = append(sortedArr, greather...)

	return sortedArr
}
