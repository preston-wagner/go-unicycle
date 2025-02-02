package slices_ext

import (
	"reflect"
	"testing"
)

func TestUnique(t *testing.T) {
	result := Unique([]int{1, 2, 3, 3, 1, 4, 5, 4})
	if !reflect.DeepEqual(result, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Unique() returned unexpected %d", result)
	}

	if !reflect.DeepEqual(Unique[int](nil), []int{}) {
		t.Error("Unique(nil) should return a slice")
	}
}
