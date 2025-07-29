package basic

import (
	"testing"
)

func TestTraverse(t *testing.T) {
	b := Basic{Items: []int{1, 2, 3, 4, 5}}
	b.TraverseFwd()
	b.TraverseBwd()
}

func TestFindMinMax(t *testing.T) {
	b := Basic{Items: []int{8, 2, 7, 3, 9}}

	if min := b.FindMin(); min != 2 {
		t.Errorf("Expected min = 2, got %d", min)
	}

	if max := b.FindMax(); max != 9 {
		t.Errorf("Expected max = 9, got %d", max)
	}
}

func TestInsertOperations(t *testing.T) {
	b := Basic{Items: []int{1, 2, 3}}

	b.InsertFirst(0)
	if b.Items[0] != 0 {
		t.Errorf("InsertFirst failed, got %v", b.Items)
	}

	b.InsertLast(4)
	if b.Items[len(b.Items)-1] != 4 {
		t.Errorf("InsertLast failed, got %v", b.Items)
	}

	b.InsertAtIndex(9, 2)
	if b.Items[2] != 9 {
		t.Errorf("InsertAtIndex failed, got %v", b.Items)
	}
}

func TestDeleteOperations(t *testing.T) {
	b := Basic{Items: []int{1, 2, 3, 4}}

	b.DeleteFirst()
	if b.Items[0] != 2 {
		t.Errorf("DeleteFirst failed, got %v", b.Items)
	}

	b.DeleteLast()
	if b.Items[len(b.Items)-1] != 3 {
		t.Errorf("DeleteLast failed, got %v", b.Items)
	}

	b.DeleteAtIndex(0)
	if len(b.Items) != 1 || b.Items[0] != 3 {
		t.Errorf("DeleteAtIndex failed, got %v", b.Items)
	}
}

func TestCountFreq(t *testing.T) {
	b := Basic{Items: []int{1, 2, 3, 2, 2, 4}}
	count := b.CountFreq(2)
	if count != 3 {
		t.Errorf("Expected CountFreq(2) = 3, got %d", count)
	}
}

func TestRotateLeftRight(t *testing.T) {
	b := Basic{Items: []int{1, 2, 3}}
	b.RotateLeft()
	expectedLeft := []int{2, 3, 1}
	for i := range expectedLeft {
		if b.Items[i] != expectedLeft[i] {
			t.Errorf("RotateLeft failed, got %v", b.Items)
			break
		}
	}

	b.RotateRight()
	expectedRight := []int{1, 2, 3}
	for i := range expectedRight {
		if b.Items[i] != expectedRight[i] {
			t.Errorf("RotateRight failed, got %v", b.Items)
			break
		}
	}
}
