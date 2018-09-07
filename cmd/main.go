/*
 *  Revision History:
 *      Initial: 2018/08/09    Wang Huajian
 */

package main

import (
	"database/sql"
	"encoding/xml"
	"fmt"

	"github.com/TechCatsLab/rss/version/v1"
	"github.com/TechCatsLab/rss/version/v2"
	"github.com/TechCatsLab/rss/client"
	"github.com/TechCatsLab/rss/model/mysql"
)

type item struct {
	version int
	url     []string
}

var gather = []item{
	item{
		version:1,
		url:[]string{
			"https://stackoverflow.com/feeds/",
			"http://www.ruanyifeng.com/blog/atom.xml",
		},
	},
	item{
		version:2,
		url:[]string{
			"http://www.adaymag.com/feed", // 译言
			"https://pansci.asia/feed", // 科学
			"https://www.echojs.com/rss", // front
		},
	},
}

func main() {
	var (
		rss1 v1.Feed
		rss2 v2.Channel
	)

	for _, gatherElement := range gather {
		switch gatherElement.version {
		case 1:
			for _, urlElement := range gatherElement.url {
				resp, err := client.Read(urlElement)
				if err != nil {
					fmt.Printf("Read from %s with error: %v\n", urlElement, err)
					return
				}
				defer resp.Close()

				decoder := xml.NewDecoder(resp)
				if err := decoder.Decode(&rss1); err != nil {
					fmt.Printf("Decode XML error: %v\n", err)
					return
				}
			}
		case 2:
			for _, urlElement := range gatherElement.url {
				resp, err := client.Read(urlElement)
				if err != nil {
					fmt.Printf("Read from %s with error: %v\n", urlElement, err)
					return
				}
				defer resp.Close()

				decoder := xml.NewDecoder(resp)
				if err := decoder.Decode(&rss2); err != nil {
					fmt.Printf("Decode XML error: %v\n", err)
					return
				}
			}
		}
	}

	// access database but no connection to the database
    db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/rss")
    if err != nil {
    	fmt.Println("open db", err)
    	return
	}
    defer db.Close()

    // assignment for StoreService
	mysql.InitStoreService(db)

	// create feed table
	mysql.StoreService.FeedServiceProvider().CreateTable()
	mysql.StoreService.FeedServiceProvider().Create(rss1.Title, rss1.Subtitle, string(rss1.Updated))
	mysql.StoreService.FeedServiceProvider().Select()

	// create entry table
	mysql.StoreService.EntryServiceProvider().CreateTable()
	for _, singleEntry := range rss1.Entry {
		mysql.StoreService.EntryServiceProvider().Create(singleEntry.Title, singleEntry.Link, string(singleEntry.Published))
	}
	mysql.StoreService.EntryServiceProvider().Select()
	//fmt.Println(rss1)

	// create channel table
	mysql.StoreService.ChannelServiceProvider().CreateTable()
	mysql.StoreService.ChannelServiceProvider().Create(rss2.Title, rss2.Description, string(rss2.LastBuildDate))
	mysql.StoreService.ChannelServiceProvider().Select()

	// crate item table
	mysql.StoreService.ItemServiceProvider().CreateTable()
	mysql.StoreService.ItemServiceProvider().Create(rss2.Title, rss2.Link, string(rss2.LastBuildDate))
	mysql.StoreService.ItemServiceProvider().Select()
}
