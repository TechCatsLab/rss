/*
 *  Revision History:
 *      Initial: 2018/08/09    Wang Huajian
 */

package main

import (
	"fmt"
	"encoding/xml"
	"github.com/TechCatsLab/rss/version/v1"
	"github.com/TechCatsLab/rss/client"
	"github.com/TechCatsLab/rss/database"
)

var (
	url1 = "https://stackoverflow.com/feeds/"
)
func main()  {
	var (
		rss1 v1.Feed
	)

	resp, err := client.Read(url1)
	if err != nil {
		fmt.Printf("Read from %s with error: %v\n", url1, err)
		return
	}
	defer resp.Close()

	decoder := xml.NewDecoder(resp)
	if err := decoder.Decode(&rss1); err != nil {
		fmt.Printf("Decode XML error: %v\n", err)
		return
	}

    database.InitMysql()

	fmt.Println(rss1)
}
