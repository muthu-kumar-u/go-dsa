package search

import (
	"math"
	"sort"
)

type Item struct {
	ID   int
	Name string
}

type Search struct {
	Items *[]Item
}

func NewSearch() *Search {
	return &Search{
		Items: &[]Item{},
	}
}

func (s *Search) LinearSearch(tId int) *Item {
	for _, i := range *s.Items {
		if i.ID == tId {
			return &i
		}
	}
	return nil
}

func (s *Search) BinarySearch(tId int) *Item {
	sorted := s.sort()
	var start = 0
	var end = len(sorted) - 1
	var mid = (start + end) / 2

	if sorted != nil {
		for start <= end  {
			if sorted[mid].ID == tId {
				res := sorted[mid]
				return &res
			} else if tId < sorted[mid].ID {
				start = mid - 1
			} else if tId > sorted[mid].ID {
				end = mid + 1
			} 
		} 
	}

	return nil
}

func (s *Search) FindFirstAndLastOccuranceLinear(tId int) (int, int) {
	sorted := s.sort()
	if sorted != nil {
		firstIndex := -1
		lastIndex := -1

		for i := range sorted {
			if sorted[i].ID == tId {
				if firstIndex == -1 {
					firstIndex = i
				}
			
				lastIndex = i
			}
		}

		return firstIndex, lastIndex
	}

	return -1, -1
}

func (s *Search) FindFirstAndLastOccuranceBinary(tId int) (int, int) {
	sorted := s.sort()
	if sorted != nil {
		firstIndex := -1
		lastIndex := -1
		
		start, end := 0, len(sorted)-1
		for start <= end {
			mid := (start + end) / 2
			if sorted[mid].ID == tId {
				firstIndex = mid
				end = mid - 1 
			} else if tId < sorted[mid].ID {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

		start, end = 0, len(sorted)-1
		for start <= end {
			mid := (start + end) / 2
			if sorted[mid].ID == tId {
				lastIndex = mid
				start = mid + 1  
			} else if tId < sorted[mid].ID {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}

		return firstIndex, lastIndex
	}

	return -1, -1
}

func (s *Search) CountOccuranceUnSorted(tId int) int {
	var count int

	if len(*s.Items) > 0 {
		for _, d := range *s.Items {
			if d.ID == tId {
				count++
			}
		}
	}

	return count
}

func (s *Search) CountOccuranceSorted(tId int) int {
	sorted := s.sort()
	var count int
	if sorted != nil {
		for _, d := range sorted {
			if d.ID == tId {
				count++
			}
		}
	}

	return count
}

// Return the index if found; otherwise return the insertion position in the sorted order.
// (Currently only inserts when the list is empty.)
func (s *Search) SearchInsertPositionBinary(tId int, name string) int {
	sorted := s.sort()

	if len(sorted) == 0 {
		*s.Items = append(*s.Items, Item{ID: tId, Name: name})
		return 0
	}

	start, end := 0, len(sorted)-1
	for start <= end {
		mid := (start + end) /2 
		if sorted[mid].ID == tId {
			return mid
		} else if sorted[mid].ID < tId {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return start
}

func (s *Search) SearchRotatedArrayBinary(tId int, arr []Item) int {
	// arr is example = [4, 5, 6, 7, 0, 1, 2] and tId is 0
	start, end := 0, len(arr)-1
	for start <= end {
		mid := (start + end) /2 
		if tId == arr[mid].ID { // 
			return mid
		} else if arr[start].ID <= arr[mid].ID { // Check if left half is sorted 
			// Check if tId lies within the sorted left half
			if arr[start].ID <= tId && tId < arr[mid].ID {
				end = mid - 1
			} else {
				start = mid + 1 // search right
			}
		} else { // Else Right half might be sorted
			if arr[mid].ID < tId && tId <= arr[end].ID {
				start = mid + 1
			} else {
				end = mid - 1 // search left
			}
		}
	}

	return -1
}

func (s *Search) FindPeakElementIndexInUnsortedUsingBinary() int { // 
	start, end := 0, len(*s.Items) - 1

	for start < end {
		mid := (start + end) / 2
		if (*s.Items)[mid].ID > (*s.Items)[mid + 1].ID {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func (s *Search) JumpSearch(tId int) int { // sorted array
	sorted := s.sort()
	jump := int(math.Floor(math.Sqrt(float64(len(sorted)))))

	var start, end int

	for i := 0; i < len(sorted); i+=jump { // O √n
		if sorted[i].ID == tId {
			return i
		} else if sorted[i].ID >= tId || i == len(sorted) -1  {
			start = i - jump
			if start < 0 {
				start = 0
			}

			end = i
			break
		}
	}

	for i := start; i <= end; i++ { // O (n)
		if sorted[i].ID == tId {
			return start
		}
	}

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