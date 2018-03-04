package sitemap

import (
	"strings"
	"testing"
)

func Test_URLTreeAppend(t *testing.T) {
	testData := []struct {
		Input  []string
		Output string
	}{
		{
			Input:  []string{"root/data", "root/data/1"},
			Output: `{"url":"root","children":[{"url":"data","children":[{"url":"1","children":[]}]}]}`,
		},
	}

	for _, tt := range testData {
		tr := NewURLTree(strings.Split(tt.Input[0], "/")[0])
		for _, data := range tt.Input {
			tr.Append(strings.Split(data, "/"))
		}

		out, err := tr.ToJSON(false)
		if err != nil {
			t.Fatal(err)
		}

		if out != tt.Output {
			t.Fatalf("unexpected output: %s != %s", tt.Output, out)
		}
	}
}
