package algos

//Bubble sort repeatedly steps through the list,
// compares adjacent elements, and swaps them if
//they are in the wrong order
//
//continues to do so until the list is sorted
//
//has a time complexity of O(n^2)
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				//swap arr[j] and arr[j+1]
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

//largest number in array pops to the end of the array with each iteration
