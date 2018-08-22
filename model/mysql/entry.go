/*
 *  Revision History:
 *      Initial: 2018/08/14    Wang Huajian
 */

package mysql

import (
	"time"
    _ "github.com/go-sql-driver/mysql"
)

const (
sqlEntryTableCreate = iota
)

type (
	Entry struct {
		id          string
		title       string
		link        string
		published   time.Time
		updated     time.Time
	}

	entryServiceProvider struct {
		store *store
	}
)

var (
	sqlEntry = []string {
	`CREATE TABLE IF NOT EXISTS entry (
		id        INT UNSIGNED AUTO_INCREMENT,
		title     VARCHAR(512) NOT NULL,
		link      VARCHAR(512) NOT NULL,
		published TIMESTAMP(6),
		PRIMARY KEY(id)
	 )ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;`,
	}
)

func (e *entryServiceProvider) CreateTable() error {
	_, err := e.store.db.Exec(sqlEntry[sqlEntryTableCreate])

	return err
}

func (e *entryServiceProvider) Create(title, link, published string) error {
	return nil
}

func (e *entryServiceProvider) Select() error {
	return nil
}
