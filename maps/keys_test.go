package maps_test

import (
	"testing"

	"github.com/preston-wagner/unicycle/maps"
	"github.com/preston-wagner/unicycle/slices"
)

func TestKeys(t *testing.T) {
	input := map[string]int{
		"a": 1,
		"b": 1,
		"c": 1,
	}
	result := maps.Keys(input)
	if len(result) != len(input) {
		t.Errorf("Keys() failed to return the correct number of keys")
	}
	if (!slices.Includes(result, "a")) || (!slices.Includes(result, "b")) || (!slices.Includes(result, "c")) {
		t.Errorf("Keys() failed to return expected values")
	}
}
