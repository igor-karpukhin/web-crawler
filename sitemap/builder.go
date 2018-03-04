package sitemap

import (
	"net/url"
	"strings"
)

func removeDuplicates(urls []string) []string {
	found := map[string]bool{}

	for _, URL := range urls {
		if _, ok := found[URL]; !ok {
			found[URL] = true
		}
	}

	var result []string
	for k := range found {
		result = append(result, k)
	}
	return result
}

func BuildJSONMap(domain string, urls []string) (string, error) {
	domainURL, err := url.Parse(domain)
	if err != nil {
		return "", err
	}
	filtered := removeDuplicates(urls)
	tree := NewURLTree(domainURL.Host)
	for _, URL := range filtered {
		u, err := url.Parse(URL)
		if err != nil {
			return "", err
		}
		tree.Append(strings.Split(u.Host+u.Path, "/"))
	}
	return tree.ToJSON(true)
}
