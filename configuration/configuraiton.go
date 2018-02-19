package configuration

import (
	"flag"
	"strings"
)

type ApplicationConfiguration struct {
	TargetURL string
	Depth     int
}

func (a *ApplicationConfiguration) Init() {
	if a == nil {
		*a = ApplicationConfiguration{}
	}

	flag.StringVar(&a.TargetURL, "target", "", "target web site to crawl")
	flag.IntVar(&a.Depth, "depth", 1, "fetching depth")

	flag.Parse()

	if strings.HasSuffix(a.TargetURL, "/") {
		a.TargetURL = strings.TrimSuffix(a.TargetURL, "/")
	}
}
