package sort

import (
	"math/rand/v2"
)

type Sort struct {}

// BubbleSort is a simple comparison-based sorting algorithm.
// It repeatedly steps through the list, compares adjacent elements,
// and swaps them if they are in the wrong order.
// Time Complexity: O(n²) in worst and average case.  Stable sort.
func (s *Sort) BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		// Last i elements are already in place
		for j := 0; j < len(arr)-i-1; j++ {
			// Compare adjacent pairs
			if arr[j] > arr[j+1] {
				// Swap if the current element is greater than next element
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}

// SelectionSort divides the array into sorted and unsorted parts.
// It repeatedly selects the smallest element from the unsorted subarray
// and puts it at its correct sorted position.
// Time Complexity: O(n²) always.  Not stable.
func (s *Sort) SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		// Assume the element at i is the smallest in remaining list
		for j := i + 1; j < len(arr); j++ {
			// Find a smaller element and swap
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

// InsertionSort builds the final sorted array one item at a time.
// It takes each element and inserts it into its correct position
// in the already sorted part of the array.
// Time Complexity: O(n²) worst, but O(n) best when nearly sorted. Stable.
func (s *Sort) InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]   // Element to be inserted
		j := i - 1

		// Shift elements in sorted part that are greater than key
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // Move element one position to the right
			j--
		}

		// Insert key in correct position
		arr[j+1] = key
	}
	return arr
}

func (s *Sort) MergeSort(arr []int) []int {
	if len(arr) <= 1 { // on single element case
		return arr
	}

	mid := len(arr) / 2
	leftHalf := arr[:mid]
	rightHalf := arr[mid:]

	s.MergeSort(leftHalf)
	s.MergeSort(rightHalf)
	s.merge(leftHalf, rightHalf, arr)

	return arr
}

func (s *Sort) merge(lH, rH, arr []int) {
	i, j, k := 0, 0, 0

	// compare elements from both halves
	for i < len(lH) && j < len(rH) {
		if lH[i] < rH[j] {
            arr[k] = lH[i]
            i++
        } else {
            arr[k] = rH[j]
            j++
        }
        k++
	}	

	// copy remaining from left half
	for i < len(lH) {
		arr[k] = lH[i]
		i++ 
		k++
	}

	// copy remaining from right half
	for j < len(rH) {
		arr[k] = rH[j]
		j++ 
		k++
	}
}

func (s *Sort) QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	randomPivot := rand.IntN(len(arr))
	arr[randomPivot], arr[len(arr)-1] = arr[len(arr)-1], arr[randomPivot]
	
	pivotIndex := s.partition(arr)
	s.QuickSort(arr[:pivotIndex])
	s.QuickSort(arr[pivotIndex+1:])

	return arr
}

func (s *Sort) partition(arr []int) int {
	pivotValue := arr[len(arr)-1]
	i, storeIndex := 0, 0

	for i < len(arr) - 1 {
		if arr[i] < pivotValue {
			arr[i], arr[storeIndex] = arr[storeIndex], arr[i]
			storeIndex++ 
		}

		i++
	}
	
	arr[storeIndex], arr[len(arr)-1] = arr[len(arr)-1], arr[storeIndex]
	return storeIndex
}