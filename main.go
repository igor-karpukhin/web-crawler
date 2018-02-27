package main

import (
	"fmt"

	"github.com/igor-karpukhin/web-crawler/configuration"
	"github.com/igor-karpukhin/web-crawler/crawler"
	"github.com/igor-karpukhin/web-crawler/sitemap"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	c := configuration.ApplicationConfiguration{}
	c.Init()
	crwl := crawler.NewCrawler(c.TargetURL, c.Depth, logger)
	crwl.Run()

	results := crwl.GetResults()
	mp, err := sitemap.BuildJSONMap(c.TargetURL, results)
	if err != nil {
		logger.Error("unable to build json map", zap.Error(err))
	}
	fmt.Println(mp)
}
