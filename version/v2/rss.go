/*
 *  Revision History:
 *      Initial: 2018/08/09    Wang Huajian
 */

package v2

import (
	"github.com/TechCatsLab/rss/version"
)

type RSS struct {
	Channels Channel `xml:"channel" json:"channel"`
}

//Channel struct for RSS
type Channel struct {
	Title         string         `xml:"title" json:"title"`
	Link          string         `xml:"link" json:"link"`
	Description   string         `xml:"description" json:"description"`
	LastBuildDate version.Date   `xml:"lastBuildDate" json:"lastBuildDate"`
	Language      string         `xml:"language" json:"language"`
	Image         []Image        `xml:"image" json:"image"`
	Item          []Item         `xml:"item" json:"item"`
}

//ItemEnclosure struct for each Item Enclosure
type ItemEnclosure struct {
	URL  string `xml:"url,attr" json:"url"`
	Type string `xml:"type,attr" json:"type"`
}

//Item struct for each Item in the Channel
type Item struct {
	Title       string          `xml:"title" json:"title"`
	Link        string          `xml:"link" json:"link"`
	Comments    string          `xml:"comments" json:"comments"`
	PubDate     version.Date    `xml:"pubDate" json:"pubDate"`
	GUID        string          `xml:"guid" json:"guid"`
	Category    []string        `xml:"category" json:"category"`
	Enclosure   []ItemEnclosure `xml:"enclosure" json:"enclosure"`
	Description string          `xml:"description" json:"description"`
	Author 		string          `xml:"author" json:"author"`
	Content     string          `xml:"content" json:"content"`
	FullText    string          `xml:"full-text" json:"full-text"`
}

type Image struct {
	Url          string          `xml:"url" json:"url"`
	Title        string	         `xml:"title" json:"title"`
	Link         string          `xml:"link" json:"link"`
}
