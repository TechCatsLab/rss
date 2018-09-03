/*
 *  Revision History:
 *      Initial: 2018/09/03    Wang Huajian
 */

package mysql

import (
	"time"
)

const (
	sqlChannelTableCreate = iota
	sqlChannelTableInsert
)

type (
	Channel struct {
		Title 			string
		Link  			string
		Description 	string
		LastBuildDate   time.Time
	}

	channelServiceProvider struct {
		store *store
	}
)
