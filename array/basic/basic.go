package basic

import (
	"log"

	constants "github.com/muthu-kumar-u/go-dsa-impl/const"
)

type Basic struct {
	Items []int
}

// Time: O(1), Space: O(1) [ignoring log output]
func (b *Basic) logItems(msg string) {
	log.Printf("%s%s: %v%s", constants.Green, msg, b.Items, constants.Reset)
}

// Time: O(n), Space: O(1)
func (b *Basic) TraverseFwd() {
	log.Println(constants.Blue + "Forward Traversal:" + constants.Reset)
	for _, d := range b.Items {
		log.Printf("  %d", d)
	}
	b.logItems("After TraverseFwd")
}

// Time: O(n), Space: O(1)
func (b *Basic) TraverseBwd() {
	log.Println(constants.Blue + "Backward Traversal:" + constants.Reset)
	for i := len(b.Items) - 1; i >= 0; i-- {
		log.Printf("  %d", b.Items[i])
	}
	b.logItems("After TraverseBwd")
}

// Time: O(n), Space: O(1)
func (b *Basic) FindMin() int {
	if len(b.Items) == 0 {
		log.Println(constants.Red + "Cannot find min in empty array" + constants.Reset)
		return -1
	}
	min := b.Items[0]
	for _, v := range b.Items {
		if v < min {
			min = v
		}
	}
	log.Printf(constants.Green+"Min value: %d"+constants.Reset, min)
	return min
}

// Time: O(n), Space: O(1)
func (b *Basic) FindMax() int {
	if len(b.Items) == 0 {
		log.Println(constants.Red + "Cannot find max in empty array" + constants.Reset)
		return -1
	}
	max := b.Items[0]
	for _, v := range b.Items {
		if v > max {
			max = v
		}
	}
	log.Printf(constants.Green+"Max value: %d"+constants.Reset, max)
	return max
}

// Time: O(1), Amortized Space: O(1)
func (b *Basic) InsertLast(v int) {
	b.Items = append(b.Items, v)
	log.Printf(constants.Blue+"Inserted %d at end"+constants.Reset, v)
	b.logItems("After InsertLast")
}

// Time: O(n), Space: O(n) [due to slice copy]
func (b *Basic) InsertFirst(v int) {
	b.Items = append([]int{v}, b.Items...)
	log.Printf(constants.Blue+"Inserted %d at start"+constants.Reset, v)
	b.logItems("After InsertFirst")
}

// Time: O(1) if replacing, O(n) if inserting (this version replaces)
func (b *Basic) InsertAtIndex(v, i int) {
	if i >= 0 && i < len(b.Items) {
		b.Items[i] = v
		log.Printf(constants.Blue+"Replaced index %d with value %d"+constants.Reset, i, v)
	} else {
		log.Println(constants.Red + "InsertAtIndex: Index out of bounds" + constants.Reset)
	}
	b.logItems("After InsertAtIndex")
}

// Time: O(1), Space: O(1)
func (b *Basic) DeleteLast() {
	if len(b.Items) == 0 {
		log.Println(constants.Red + "DeleteLast: Empty array" + constants.Reset)
		return
	}
	b.Items = b.Items[:len(b.Items)-1]
	log.Println(constants.Blue + "Deleted last item" + constants.Reset)
	b.logItems("After DeleteLast")
}

// Time: O(n), Space: O(n) [due to slice reallocation]
func (b *Basic) DeleteFirst() {
	if len(b.Items) == 0 {
		log.Println(constants.Red + "DeleteFirst: Empty array" + constants.Reset)
		return
	}
	b.Items = b.Items[1:]
	log.Println(constants.Blue + "Deleted first item" + constants.Reset)
	b.logItems("After DeleteFirst")
}

// Time: O(n), Space: O(n) [due to slice reallocation]
func (b *Basic) DeleteAtIndex(i int) {
	if i >= 0 && i < len(b.Items) {
		b.Items = append(b.Items[:i], b.Items[i+1:]...)
		log.Printf(constants.Blue+"Deleted value at index %d"+constants.Reset, i)
	} else {
		log.Println(constants.Red + "DeleteAtIndex: Index out of bounds" + constants.Reset)
	}
	b.logItems("After DeleteAtIndex")
}

// Time: O(n), Space: O(1)
func (b *Basic) CountFreq(v int) int {
	count := 0
	for _, val := range b.Items {
		if val == v {
			count++
		}
	}
	log.Printf(constants.Green+"Count of %d: %d"+constants.Reset, v, count)
	return count
}

// Time: O(n), Space: O(n)
func (b *Basic) RotateLeft() {
	if len(b.Items) <= 1 {
		log.Println(constants.Red + "RotateLeft: Not enough elements" + constants.Reset)
		return
	}
	first := b.Items[0]
	b.Items = append(b.Items[1:], first)
	log.Println(constants.Blue + "Rotated array to the left" + constants.Reset)
	b.logItems("After RotateLeft")
}

// Time: O(n), Space: O(n)
func (b *Basic) RotateRight() {
	if len(b.Items) <= 1 {
		log.Println(constants.Red + "RotateRight: Not enough elements" + constants.Reset)
		return
	}
	last := b.Items[len(b.Items)-1]
	b.Items = append([]int{last}, b.Items[:len(b.Items)-1]...)
	log.Println(constants.Blue + "Rotated array to the right" + constants.Reset)
	b.logItems("After RotateRight")
}
