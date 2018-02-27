package sitemap

import (
	"encoding/json"
	"fmt"
)

type URLTree struct {
	URL   string     `json:"url"`
	Nodes []*URLTree `json:"children"`
}

func NewURLTree(URL string) *URLTree {
	return &URLTree{
		Nodes: []*URLTree{},
		URL:   URL,
	}
}

func (t *URLTree) Append(urls []string) {
	if len(urls)-1 == 0 {
		return
	}
	if t.URL == urls[0] {
		childFound := false
		for i := 0; i < len(t.Nodes); i++ {
			if t.Nodes[i].URL == urls[1] {
				childFound = true
				t.Nodes[i].Append(urls[1:])
			}

		}
		if !childFound {
			tree := NewURLTree(urls[1])
			t.Nodes = append(t.Nodes, tree)
			tree.Append(urls[1:])
		}
	}
}

func (t *URLTree) ToJSON() (string, error) {
	r, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "", err
	}
	return string(r), nil
}

func (t *URLTree) PrintTree(indent string) {
	if t != nil {
		fmt.Printf("%s%s\r\n", indent, t.URL)
		for i := 0; i < len(t.Nodes); i++ {
			t.Nodes[i].PrintTree(indent + "  ")
		}
	}
}
