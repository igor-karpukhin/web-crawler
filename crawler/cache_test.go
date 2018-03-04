package crawler

import (
	"reflect"
	"testing"
)

func Test_URLCacheAdd(t *testing.T) {
	testData := []struct {
		Input  []string
		Output map[string]bool
	}{
		{
			Input: []string{"a", "b", "c"},
			Output: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
			},
		},
	}

	cache := NewURLCache()

	for _, tt := range testData {
		for _, data := range tt.Input {
			cache.Add(data)
		}

		out := cache.FetchAll()
		if !reflect.DeepEqual(tt.Output, out) {
			t.Fatalf("Output is not equal to expected: %v != %v",
				tt.Output, out)
		}
	}
}

func Test_URLCacheHas(t *testing.T) {
	testData := []struct {
		Input  []string
		Output bool
	}{
		{
			Input:  []string{"a", "b", "c"},
			Output: true,
		},
	}

	cache := NewURLCache()

	for _, tt := range testData {
		for _, data := range tt.Input {
			cache.Add(data)
			out := cache.Has(data)
			if !out {
				t.Fatalf("Output is not equal to expected: %v != %v",
					tt.Output, out)
			}
		}
	}
}
