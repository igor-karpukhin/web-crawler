package configuration

import (
	"flag"
	"strings"
)

type ApplicationConfiguration struct {
	TargetURL  string
	Depth      int
	Version    bool
	OutputFile string
}

func (a *ApplicationConfiguration) Init() {
	if a == nil {
		*a = ApplicationConfiguration{}
	}

	flag.StringVar(&a.TargetURL, "target", "", "target web site to crawl")
	flag.IntVar(&a.Depth, "depth", 1, "fetching depth")
	flag.BoolVar(&a.Version, "v", false, "print version and exit application")
	flag.StringVar(&a.OutputFile, "o", "", "output file for site tree")

	flag.Parse()

	if strings.HasSuffix(a.TargetURL, "/") {
		a.TargetURL = strings.TrimSuffix(a.TargetURL, "/")
	}
}
