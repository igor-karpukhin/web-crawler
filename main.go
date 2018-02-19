package main

import (
	"github.com/igor-karpukhin/web-crawler/configuration"
	"github.com/igor-karpukhin/web-crawler/crawler"
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
}
