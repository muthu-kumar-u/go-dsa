package main

import "github.com/muthu-kumar-u/go-dsa-impl/array/basic"

func main() {
	b := basic.Basic{
		Items: []int{5, 2, 9, 1, 3},
	}

	b.TraverseFwd()
	b.TraverseBwd()

	b.FindMin()
	b.FindMax()

	b.InsertFirst(10)
	b.InsertLast(20)
	b.InsertAtIndex(99, 2)

	b.DeleteFirst()
	b.DeleteLast()
	b.DeleteAtIndex(1)

	b.CountFreq(9)

	b.RotateLeft()
	b.RotateRight()
}