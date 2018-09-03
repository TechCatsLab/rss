/*
 *  Revision History:
 *      Initial: 2018/09/03    Wang Huajian
 */

package mysql

import (
	"time"
	"errors"
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

var (
	errChannelInvalidInsert = errors.New("insert: insert affected 0 rows")
	sqlChannel = []string {
		`CREATE TABLE IF NOT EXISTS feed (
            channel_id   INTEGER UNSIGNED AUTO_INCREMENT, 
            title     VARCHAR(512) NOT NULL,
            description VARCHAR(512) NOT NUNLL,
            lastbuilddate   VARCHAR(128),
			INDEX(title),
            PRIMARY KEY(channel_id)
         ) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
		`INSERT INTO channel(title, description,lastbuilddate) VALUES (?,?,?)`,
	}
)

// create channel table
func (c *channelServiceProvider) CreateTable() error {
	_, err := c.store.db.Exec(sqlChannel[sqlChannelTableCreate])

	return err
}

// insert a raw
func (c *channelServiceProvider) Create(title, desc, last string) error {
	result, err := c.store.db.Exec(sqlChannel[sqlChannelTableInsert], title, desc, last)
	if err != nil {
		return err
	}
	if affected, _ := result.RowsAffected(); affected == 0 {
		return  errChannelInvalidInsert
	}
	return nil
}

func (c *channelServiceProvider) Select() error {
	return nil
}

func (c *channelServiceProvider) Delete() error {
	return nil
}
