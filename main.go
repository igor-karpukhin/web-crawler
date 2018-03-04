package main

import (
	"fmt"

	"os"

	"github.com/igor-karpukhin/web-crawler/configuration"
	"github.com/igor-karpukhin/web-crawler/crawler"
	"github.com/igor-karpukhin/web-crawler/dataprovider"
	"github.com/igor-karpukhin/web-crawler/sitemap"
	"github.com/igor-karpukhin/web-crawler/version"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	c := configuration.ApplicationConfiguration{}
	c.Init()

	if c.Version {
		fmt.Printf("Version: %s; Build time %s\r\n", version.Version,
			version.BuildTime)
		return
	}

	provider := dataprovider.NewHTMLDataProvider()

	crwl := crawler.NewCrawler(provider, c.TargetURL, c.Depth, logger)
	crwl.Run()

	results := crwl.GetResults()
	mp, err := sitemap.BuildJSONMap(c.TargetURL, results)
	if err != nil {
		logger.Error("unable to build json map", zap.Error(err))
	}
	if c.OutputFile != "" {
		logger.Info("Writing map into " + c.OutputFile)
		f, err := os.OpenFile(c.OutputFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			logger.Error("unable to open file", zap.Error(err))
			return
		}
		n, err := f.WriteString(mp)
		if err != nil {
			logger.Error("unable to store map", zap.Error(err))
			return
		}
		logger.Info(fmt.Sprintf("Done. %d bytes written", n))

	} else {
		fmt.Println(mp)
	}
}
