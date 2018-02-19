package crawler

import (
	"net/http"

	"strconv"
	"strings"
	"sync"

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

func NewCrawler(domainUrl string, depth int, logger *zap.Logger) *Crawler {
	return &Crawler{
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

	resp, err := http.Get(rootUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}

	root := html.NewTokenizer(resp.Body)
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
					c.l.Info("URL: " + url + "; Depth: " + strconv.Itoa(c.depth-depth+1))
					wg.Add(1)
					go c.work(wg, url, depth-1)
				}
			}
		}
	}
}

func (c *Crawler) Run() error {
	if c.running {
		return errors.New("already running")
	}

	c.l.Info("started to crawl", zap.String("domain", c.domainUrl))
	c.wg.Add(1)
	c.work(c.wg, c.domainUrl, c.depth)
	c.wg.Wait()
	c.l.Info("Done")
	c.l.Info("results", zap.Any("r", c.urlCache.FetchAll()))
	return nil
}
