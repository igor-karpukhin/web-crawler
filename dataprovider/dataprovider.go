package dataprovider

import "io"

type DataProvider interface {
	Fetch(url string) (io.Reader, error)
}
