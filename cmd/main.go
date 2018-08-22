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
	"github.com/TechCatsLab/rss/client"
	"github.com/TechCatsLab/rss/model/mysql"
)

var (
	urlStackoverflow = "https://stackoverflow.com/feeds/"
)

func main()  {
	var (
		rss1 v1.Feed
	)

	resp, err := client.Read(urlStackoverflow)
	if err != nil {
		fmt.Printf("Read from %s with error: %v\n", urlStackoverflow, err)
		return
	}
	defer resp.Close()

	decoder := xml.NewDecoder(resp)
	if err := decoder.Decode(&rss1); err != nil {
		fmt.Printf("Decode XML error: %v\n", err)
		return
	}
	// access database but no connection to the database
    db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/rss")
    defer db.Close()

    // assignment for StoreService
	mysql.InitStoreService(db)

	// create feed table
	mysql.StoreService.FeedServiceProvider().CreateTable()

	// create entry table
	mysql.StoreService.EntryServiceProvider().CreateTable()
	fmt.Println(rss1)
}
