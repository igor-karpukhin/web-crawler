package crawler

import (
	"reflect"
	"testing"

	"github.com/igor-karpukhin/web-crawler/dataprovider"
	"go.uber.org/zap"
)

func Test_CrawlerDo(t *testing.T) {
	testData := []struct {
		RootURL      string
		DataProvider dataprovider.DataProvider
		Output       []string
	}{
		{
			RootURL: "test.com",
			DataProvider: &dataprovider.MockDataProvider{
				Data: map[string]string{
					"test.com":      `<a href="test.com/data"></a>"`,
					"test.com/data": `<a href="/forum">`,
				},
			},
			Output: []string{"test.com/data", "test.com/forum"},
		},
	}

	for _, tt := range testData {
		c := NewCrawler(tt.DataProvider, tt.RootURL, 2, zap.L())
		err := c.Run()
		if err != nil {
			t.Fatal(err)
		}
		res := c.GetResults()
		if !reflect.DeepEqual(res, tt.Output) {
			t.Fatalf("unexpected output. %v != %v", res, tt.Output)
		}
	}
}
