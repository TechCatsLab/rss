/*
 *  Revision History:
 *      Initial: 2018/08/10    Wang Huajian
 */

package mysql

import (
	"time"
	_ "github.com/go-sql-driver/mysql"
)

const (
    sqlFeedTableCreate = iota
)

type (
	Feed struct {
		Title         string
		Link          string
		Subtitle      string
		Updated       time.Time
	}

	feedServiceProvider struct {
		store *store
	}
)

var (
	sqlFeed = []string {
		`CREATE TABLE IF NOT EXISTS feed (
            feed_id   INTEGER UNSIGNED AUTO_INCREMENT, 
            title     VARCHAR(512) NOT NULL,
            subtitle  VARCHAR(512) NOT NULL,
            updated   TIMESTAMP(6),
			INDEX(title)
         ) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8_bin;`,
	}
)

// create feed table
func (f *feedServiceProvider) CreateTable() error {
	_, err := f.store.db.Exec(sqlFeed[sqlFeedTableCreate])

	return err
}

// insert a raw
func (f *feedServiceProvider) Create(id uint32, title, subtitle string, up string) error {
	return nil
}

func (f *feedServiceProvider) Select() error {
	return nil
}