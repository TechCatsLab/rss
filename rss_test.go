/*
 *  Revision History:
 *      Initial: 2018/08/02    Wang Huajian
 */

package rss

import (
	"net/http"
	"testing"
)

var address = []string{
	"http://blog.sina.com.cn/rss/1219548027.xml",
	"http://36kr.com/feed",
	"http://www.ifanr.com/feed",
	"http://www.phonekr.com/feed/",
	"https://cn.engadget.com/rss.xml",
	}

// testFetcher is an implementation of the Fetcher interface which reads the
// content from a local file.
type testFetcher struct{}

func (f *testFetcher) Get(url string) (resp *http.Response, err error) {
	file, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func TestAllFeedsParse(t *testing.T) {
	for _, url := range address {
		if _, err := ReadWithClient(url, new(testFetcher)); err != nil {
			t.Fatalf("ReadWithClient(%q) err = %v, expected nil", url, err)
		}
	}
}
