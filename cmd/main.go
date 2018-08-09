/*
 *  Revision History:
 *      Initial: 2018/08/09    Wang Huajian
 */

package main

import (
	"fmt"
	"encoding/xml"
	"github.com/TechCatsLab/rss/client"
	"github.com/TechCatsLab/rss/version/v2"
)

var (
	url = "http://www.geekpark.net/rss"
)

func main() {
	var (
		rss v2.RSS
	)

	resp, err := client.Read(url)
	if err != nil {
		fmt.Printf("Read from %s with error: %v\n", url, err)
		return
	}
	defer resp.Close()


	decoder := xml.NewDecoder(resp)
	if err := decoder.Decode(&rss); err != nil {
		fmt.Printf("Decode XML error: %v\n", err)
		return
	}

	fmt.Println(rss)
}

