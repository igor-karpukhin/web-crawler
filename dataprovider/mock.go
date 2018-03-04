package dataprovider

import (
	"errors"
	"io"
	"strings"
)

type MockDataProvider struct {
	Data map[string]string
}

func NewMockDataProvider() *MockDataProvider {
	return &MockDataProvider{
		Data: map[string]string{},
	}
}

func (m *MockDataProvider) Fetch(url string) (io.Reader, error) {
	if val, ok := m.Data[url]; ok {
		return strings.NewReader(val), nil
	} else {
		return nil, errors.New("not found")
	}
}
