/*
 *  Revision History:
 *      Initial: 2018/08/02    Wang Huajian
 */

//Package rss provides a Simple RSS parser, tested with various feeds.
package rss

import (
"encoding/xml"
"net/http"
"time"
"crypto/tls"

"github.com/paulrosania/go-charset/charset"
_ "github.com/paulrosania/go-charset/data" //initialize only
	"fmt"
)

const (
	wordpressDateFormat = "Mon, 02 Jan 2006 15:04:05 -0700"
)

//Channel struct for RSS
type Channel struct {
	Title         string `xml:"title"`
	Link          string `xml:"link"`
	Description   string `xml:"description"`
	Language      string `xml:"language"`
	LastBuildDate Date   `xml:"lastBuildDate"`
	Item          []Item `xml:"item"`
}

//ItemEnclosure struct for each Item Enclosure
type ItemEnclosure struct {
	URL  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

//Item struct for each Item in the Channel
type Item struct {
	Title       string          `xml:"title" json:"title"`
	Link        string          `xml:"link"`
	Comments    string          `xml:"comments"`
	PubDate     Date            `xml:"pubDate"`
	GUID        string          `xml:"guid"`
	Category    []string        `xml:"category"`
	Enclosure   []ItemEnclosure `xml:"enclosure"`
	Description string          `xml:"description"`
	Author 		string          `xml:"author"`
	Content     string          `xml:"content"`
	FullText    string          `xml:"full-text"`
}

//Date type
type Date string

//Parse (Date function) and returns Time, error
func (d Date) Parse() (time.Time, error) {
	t, err := d.ParseWithFormat(wordpressDateFormat)
	if err != nil {
		t, err = d.ParseWithFormat(time.RFC822) // RSS 2.0 spec
	}
	return t, err
}

//ParseWithFormat (Date function), takes a string and returns Time, error
func (d Date) ParseWithFormat(format string) (time.Time, error) {
	return time.Parse(format, string(d))
}

//Format (Date function), takes a string and returns string, error
func (d Date) Format(format string) (string, error) {
	t, err := d.Parse()
	if err != nil {
		return "", err
	}
	return t.Format(format), nil
}

//MustFormat (Date function), take a string and returns string
func (d Date) MustFormat(format string) string {
	s, err := d.Format(format)
	if err != nil {
		return err.Error()
	}
	return s
}

