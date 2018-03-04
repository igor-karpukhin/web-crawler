package dataprovider

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HTMLDataProviderFetch(t *testing.T) {
	testData := []struct {
		Input      string
		HFunc      http.HandlerFunc
		ShouldFail bool
	}{
		{
			Input: "/failed",
			HFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			},
			ShouldFail: true,
		},
		{
			Input: "/passed",
			HFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				io.WriteString(w, "<html>Test</html>")
			},
			ShouldFail: false,
		},
	}

	provider := NewHTMLDataProvider()

	for _, tt := range testData {
		serv := httptest.NewServer(tt.HFunc)
		_, err := provider.Fetch(serv.URL + tt.Input)
		serv.Close()
		if err != nil {
			if !tt.ShouldFail {
				t.Fatalf("test failed but it shouldn't. Err: %s", err)
			}
		} else {
			if tt.ShouldFail {
				t.Fatalf("test failed but it shouldn't. Err: %s", err)
			}
		}
	}
}
