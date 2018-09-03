/*
 *  Revision History:
 *      Initial: 2018/09/03    Wang Huajian
 */

package mysql

import (
	"errors"
	"time"
)

const (
	sqlItemTableCreate = iota
	sqlItemTableInsert
)

type (
	Item struct {
		Title        string
		Link         string
		PubDate      time.Time
	}

	itemServiceProvider struct {
		store *store
	}
)

var (
	errItemInvalidInsert = errors.New("insert: insert affected 0 rows")
	sqlItem = []string {
		`CREATE TABLE IF NOT EXISTS entry (
			title     VARCHAR(512) NOT NULL,
			link      VARCHAR(512) NOT NULL,
			pubdate   VARCHAR(128),
		 )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
		`INSERT INFO item(title, link, pubdate) VALUES (?,?,?)`,
	}
)

// crate item table
func (i *itemServiceProvider) CrateTable() error {
	_, err := i.store.db.Exec(sqlItem[sqlItemTableCreate])

	return err
}

// insert a raw
func (i *itemServiceProvider) Crate(title, link, pubdate string) error {
	result, err := i.store.db.Exec(sqlItem[sqlItemTableInsert], title, link, pubdate)
	if err != nil {
		return err
	}
	if affected, _ := result.RowsAffected(); affected == 0 {
		return  errItemInvalidInsert
	}
	return nil
}

func (i *itemServiceProvider) Select() error {
	return nil
}

func (i *itemServiceProvider) Delete() error {
	return nil
}
