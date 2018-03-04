package sitemap

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	testData := []struct {
		Input  []string
		Output []string
	}{
		{
			Input:  []string{"1", "2", "2", "3"},
			Output: []string{"1", "2", "3"},
		},
	}

	for _, tt := range testData {
		out := removeDuplicates(tt.Input)
		if !reflect.DeepEqual(tt.Output, out) {
			t.Fatalf("unexpected output: %v != %v", tt.Output, out)
		}
	}
}
