/*
 *  Revision History:
 *      Initial: 2018/08/09    Wang Huajian
 */

package v1

import (
	"github.com/TechCatsLab/rss/version"
)

// Feed struct for RSS version-1
type Feed struct {
	Title         string         `xml:"title" json:"title"`
	Link          string         `xml:"link" json:"link"`
	Subtitle      string         `xml:"subtitle" json:"subtitle"`
	Updated       version.Date   `xml:"updated" json:"updated"`
	Id            string         `xml:"id" json:"id"`
	Entry         []Entry        `xml:"entry" json:"entry"`
}

type Entry struct {
	Id            string         `xml:"id" json:"id"`
	Title         string         `xml:"title" json:"title"`
	Category      string         `xml:"category" json:"category"`
	Author        []Author       `xml:"author" json:"author"`
	Link          string         `xml:"link" json:"link"`
	Published     version.Date   `xml:"published" json:"published"`
	Updated       version.Date   `xml:"updated" json:"updated"`
	Summary       string         `xml:"summary" json:"summary"`
}

type Author struct {
	Name          string          `xml:"name" json:"name"`
	Uri           string          `xml:"uri" json:"uri"`
}
