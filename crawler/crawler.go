package crawler

import (
	"strings"
	"sync"

	"github.com/igor-karpukhin/web-crawler/dataprovider"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/net/html"
)

type Crawler struct {
	urlCache  *URLCache
	domainUrl string
	running   bool
	depth     int
	l         *zap.Logger
	wg        *sync.WaitGroup
	provider  dataprovider.DataProvider
}

func HasHref(token html.Token) (bool, string) {
	found := false
	var url string
	for _, attr := range token.Attr {
		if attr.Key == "href" {
			url = attr.Val
			found = true
			break
		}
	}

	return found, url
}

func NewCrawler(provider dataprovider.DataProvider, domainUrl string,
	depth int, logger *zap.Logger) *Crawler {
	return &Crawler{
		provider:  provider,
		urlCache:  NewURLCache(),
		domainUrl: domainUrl,
		depth:     depth,
		l:         logger,
		wg:        &sync.WaitGroup{},
	}
}

func (c *Crawler) work(wg *sync.WaitGroup, rootUrl string, depth int) {
	defer wg.Done()
	if depth <= 0 {
		return
	}

	resp, err := c.provider.Fetch(rootUrl)
	if err != nil {
		return
	}

	root := html.NewTokenizer(resp)
	for {
		token := root.Next()
		switch {
		case token == html.ErrorToken:
			return
		case token == html.StartTagToken:
			tt := root.Token()

			//Not a link
			if tt.Data != "a" {
				continue
			}

			//No 'href' attribute
			ok, url := HasHref(tt)
			if !ok {
				continue
			}

			//Transform relative domain link to direct link
			if strings.HasPrefix(url, "/") {
				url = c.domainUrl + url
			}

			//If domain link
			if strings.HasPrefix(url, c.domainUrl) {
				if !c.urlCache.Has(url) {
					c.urlCache.Add(url)
					wg.Add(1)
					go c.work(wg, url, depth-1)
				}
			}
		}
	}
}

func (c *Crawler) GetResults() []string {
	var result []string
	for URL, _ := range c.urlCache.FetchAll() {
		result = append(result, URL)
	}
	return result
}

func (c *Crawler) Run() error {
	if c.running {
		return errors.New("already running")
	}

	c.l.Info("started to crawl", zap.String("domain", c.domainUrl))
	c.wg.Add(1)
	c.work(c.wg, c.domainUrl, c.depth)
	c.wg.Wait()
	c.l.Info("done")
	c.l.Debug("results", zap.Any("r", c.urlCache.FetchAll()))
	return nil
}
