package search

import (
	"math"
	"sort"
)

// Item defines a simple structure with ID and Name.
type Item struct {
	ID   int
	Name string
}

// Search wraps a pointer to a slice of Items for search operations.
type Search struct {
	Items *[]Item
}

// NewSearch initializes and returns a new Search instance.
func NewSearch() *Search {
	return &Search{
		Items: &[]Item{},
	}
}

// LinearSearch performs a linear search to find the first item with the given ID.
// Time Complexity: O(n)
// Space Complexity: O(1)
func (s *Search) LinearSearch(tId int) *Item {
	for _, i := range *s.Items {
		if i.ID == tId {
			return &i // Return first matching item
		}
	}
	return nil // Not found
}

// BinarySearch performs a classic binary search on a sorted array to find an item by ID.
// Returns a pointer to the matched item if found, otherwise nil.
// Time Complexity: O(log n)
// Space Complexity: O(1)
func (s *Search) BinarySearch(tId int) *Item {
	sorted := s.sort()
	start := 0
	end := len(sorted) - 1

	// Binary search loop
	for start <= end {
		mid := (start + end) / 2
		if sorted[mid].ID == tId {
			res := sorted[mid] // Avoid returning pointer to loop variable
			return &res
		} else if tId < sorted[mid].ID {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return nil // Not found
}

// FindFirstAndLastOccuranceLinear finds the first and last index of tId using linear search.
// Time Complexity: O(n)
// Space Complexity: O(1)
// Returns (-1, -1) if the ID is not found.
func (s *Search) FindFirstAndLastOccuranceLinear(tId int) (int, int) {
	sorted := s.sort()

	firstIndex := -1
	lastIndex := -1

	for i := range sorted {
		if sorted[i].ID == tId {
			if firstIndex == -1 {
				firstIndex = i // First time match found
			}
			lastIndex = i // Keep updating until last match
		}
	}

	return firstIndex, lastIndex
}

// FindFirstAndLastOccuranceBinary finds the first and last index of tId using two binary searches.
// First search narrows to the leftmost occurrence.
// Second search narrows to the rightmost occurrence.
// Time Complexity: O(log n)
// Space Complexity: O(1)
// Returns (-1, -1) if ID is not found.
func (s *Search) FindFirstAndLastOccuranceBinary(tId int) (int, int) {
	sorted := s.sort()
	firstIndex := -1
	lastIndex := -1

	// Binary search for the first occurrence (leftmost)
	start, end := 0, len(sorted)-1
	for start <= end {
		mid := (start + end) / 2
		if sorted[mid].ID == tId {
			firstIndex = mid
			end = mid - 1 // Keep going left
		} else if tId < sorted[mid].ID {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	// Binary search for the last occurrence (rightmost)
	start, end = 0, len(sorted)-1
	for start <= end {
		mid := (start + end) / 2
		if sorted[mid].ID == tId {
			lastIndex = mid
			start = mid + 1 // Keep going right
		} else if tId < sorted[mid].ID {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return firstIndex, lastIndex
}

// CountOccuranceUnSorted counts how many times a given ID appears in an unsorted array.
// Time Complexity: O(n)
// Use Case: Any arbitrary list (no sorting needed)
func (s *Search) CountOccuranceUnSorted(tId int) int {
	var count int

	if len(*s.Items) > 0 {
		for _, d := range *s.Items {
			if d.ID == tId {
				count++ // Increment if ID matches
			}
		}
	}

	return count
}

// CountOccuranceSorted counts how many times a given ID appears in a sorted array.
// Time Complexity: O(n) — can be optimized to O(log n) using binary search boundaries
// Use Case: When the list is already or expected to be sorted
func (s *Search) CountOccuranceSorted(tId int) int {
	sorted := s.sort()
	var count int

	if sorted != nil {
		for _, d := range sorted {
			if d.ID == tId {
				count++ // Match found
			}
		}
	}

	return count
}

// SearchInsertPositionBinary returns the index of tId if found,
// otherwise returns the index where it should be inserted to keep the array sorted.
// Also inserts the item if the list is empty.
// Time Complexity: O(log n)
// Use Case: Used for efficient insertions in ordered structures
func (s *Search) SearchInsertPositionBinary(tId int, name string) int {
	sorted := s.sort()

	// Special case: insert into empty list
	if len(sorted) == 0 {
		*s.Items = append(*s.Items, Item{ID: tId, Name: name})
		return 0
	}

	start, end := 0, len(sorted) - 1

	// Standard binary search to find position
	for start <= end {
		mid := (start + end) / 2
		if sorted[mid].ID == tId {
			return mid // ID exists at this index
		} else if sorted[mid].ID < tId {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	// ID not found, start now represents the correct insert position
	return start
}

// SearchRotatedArrayBinary performs binary search on a rotated sorted array.
// Example input: [4, 5, 6, 7, 0, 1, 2], target: 0 → returns index 4
// Time Complexity: O(log n)
// Use Case: Arrays that are sorted but rotated at some unknown pivot
func (s *Search) SearchRotatedArrayBinary(tId int, arr []Item) int {
	start, end := 0, len(arr) - 1

	for start <= end {
		mid := (start + end) / 2

		if tId == arr[mid].ID {
			return mid // Target found
		}

		// Determine which side is properly sorted
		if arr[start].ID <= arr[mid].ID {
			// Left half is sorted
			if arr[start].ID <= tId && tId < arr[mid].ID {
				end = mid - 1 // Search left
			} else {
				start = mid + 1 // Search right
			}
		} else {
			// Right half is sorted
			if arr[mid].ID < tId && tId <= arr[end].ID {
				start = mid + 1 // Search right
			} else {
				end = mid - 1 // Search left
			}
		}
	}

	// Target not found
	return -1
}

// FindPeakElementIndexInUnsortedUsingBinary finds the index of a peak element
// in an unsorted array using a binary search approach.
// A peak is an element that is strictly greater than its right neighbor.
// Time Complexity: O(log n)
// Space Complexity: O(1)
// Note: Assumes there is at least one peak (always true unless array is empty or flat).
func (s *Search) FindPeakElementIndexInUnsortedUsingBinary() int {
	start, end := 0, len(*s.Items) - 1

	// Modified binary search:
	// - If current element is greater than the next one, peak lies on the left (including mid)
	// - Else, it lies to the right (excluding mid)
	for start < end {
		mid := (start + end) / 2
		if (*s.Items)[mid].ID > (*s.Items)[mid + 1].ID {
			end = mid // Peak is on left side
		} else {
			start = mid + 1 // Peak is on right side
		}
	}

	// When start == end, it points to the peak index
	return start
}


// JumpSearch performs search in a sorted array using jump search technique.
// It first jumps ahead by sqrt(n) steps to find the block where the target may exist,
// then does a linear scan within that block.
// Time Complexity:
//   - Jump phase: O(√n)
//   - Linear scan: O(√n) worst case
//   - Total: O(√n)
// Space Complexity: O(1)
func (s *Search) JumpSearch(tId int) int {
	sorted := s.sort() // Ensure array is sorted

	// Step size for jump: sqrt of array length
	jump := int(math.Floor(math.Sqrt(float64(len(sorted)))))

	var start, end int

	// Jump phase: find the block where tId may be present
	for i := 0; i < len(sorted); i += jump {
		if sorted[i].ID == tId {
			return i // Direct hit at jump position
		} else if sorted[i].ID >= tId || i == len(sorted)-1 {
			// Found a block that might contain tId
			start = i - jump
			if start < 0 {
				start = 0 // Ensure start doesn't go negative
			}
			end = i
			break
		}
	}

	// Linear search in the block [start, end]
	for i := start; i <= end && i < len(sorted); i++ {
		if sorted[i].ID == tId {
			return i
		}
	}

	// Target not found
	return -1
}


// ExponentialSearch performs a search on a sorted array of Items using exponential search.
// It first finds a suitable range by exponentially increasing the index,
// then applies binary search within that range to locate the target ID.
// Time Complexity: O(log i), where i is the index where target lies or should lie.
func (s *Search) ExponentialSearch(tId int) int {
	sorted := s.sort() // Ensure the array is sorted

	var start, end int

	// Base case: check the very first element
	if sorted[0].ID == tId {
		return start // return 0
	}

	var index = 1

	// Step 1: Exponential range discovery
	// Double the index each time to quickly find a potential range where tId may exist
	for i := index; i <= len(sorted)-1; i += index {
		if sorted[i].ID == tId {
			return i // If found directly at index
		} else if tId < sorted[i].ID || i == len(sorted)-1 {
			// If we've passed the target or reached end, define search window
			start = index / 2
			end = i
			break
		} else {
			index = index * 2 // Exponentially increase the step
		}
	}

	// Safety check: ensure end is within bounds
	if end >= len(sorted) {
		end = len(sorted) - 1
	}

	// Step 2: Binary search within the narrowed range
	for start <= end {
		mid := (start + end) / 2

		if sorted[mid].ID == tId {
			return mid // Target found
		} else if sorted[mid].ID > tId {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	// Target not found
	return -1
}

// TernerySearch performs a search on a sorted array of Items using ternary search.
// It divides the array into three parts by choosing two mid points,
// and narrows the search window based on where the target might lie.
// Time Complexity: O(log₃ n)
// Use Case: Works well on sorted arrays (especially unimodal functions, though not required here).
func (s *Search) TernerySearch(tId int) int {
	sorted := s.sort() // Ensure the array is sorted

	start, end := 0, len(sorted)-1

	// Loop until the search range is valid
	for start <= end {
		// Calculate the two mid points that divide the range into three parts
		mid1 := start + (end - start) / 3
		mid2 := end - (end - start) / 3

		// Check if either mid point is the target
		if sorted[mid1].ID == tId {
			return mid1
		} else if sorted[mid2].ID == tId {
			return mid2
		}

		// Narrow the search to the middle section
		if tId > sorted[mid1].ID && tId < sorted[mid2].ID {
			start = mid1 + 1
			end = mid2 - 1
		} else if tId < sorted[mid1].ID {
			// Target is in the left third
			end = mid1 - 1
		} else {
			// Target is in the right third
			start = mid2 + 1
		}
	}

	// Target not found
	return -1
}


func (s *Search) sort() []Item {
	if len(*s.Items) == 0 {
		return nil
	}
	sorted := make([]Item, len(*s.Items))
	copy(sorted, *s.Items)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].ID < sorted[j].ID
	})
	return sorted
}