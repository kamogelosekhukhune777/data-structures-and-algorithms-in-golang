package algos

//InsertionSort
//
//has a time complexity of O(n^2)
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		/* move elements of arr[0..i-1]that are grather than
		key to one position ahead of their current position*/
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
