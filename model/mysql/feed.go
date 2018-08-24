/*
 *  Revision History:
 *      Initial: 2018/08/10    Wang Huajian
 */

package mysql

import (
	"errors"
	"time"
	_ "github.com/go-sql-driver/mysql"
)

const (
    sqlFeedTableCreate = iota
    sqlFeedTableInsert
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
	errFeedInvalidInsert = errors.New("insert: insert affected 0 rows")
	sqlFeed = []string {
		`CREATE TABLE IF NOT EXISTS feed (
            feed_id   INTEGER UNSIGNED AUTO_INCREMENT, 
            title     VARCHAR(512) NOT NULL,
            subtitle  VARCHAR(512) NOT NULL,
            updated   VARCHAR(128),
			INDEX(title),
            PRIMARY KEY(feed_id)
         ) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
         "INSERT feed SET title=?,subtitle=?,updated=?",
	}
)

// create feed table
func (f *feedServiceProvider) CreateTable() error {
	_, err := f.store.db.Exec(sqlFeed[sqlFeedTableCreate])

	return err
}

// insert a raw
func (f *feedServiceProvider) Create(title, subtitle, up string) error {
	result, err := f.store.db.Exec(sqlFeed[sqlFeedTableInsert], title, subtitle, up)
	if err != nil {
		return err
	}
    if affected, _ := result.RowsAffected(); affected == 0 {
    	return  errFeedInvalidInsert
	}
	return nil
}

func (f *feedServiceProvider) Select() error {
	return nil
}

func (f *feedServiceProvider) Delete()  error {
	return nil
}
