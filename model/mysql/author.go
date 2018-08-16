/*
 *  Revision History:
 *      Initial: 2018/08/14    Wang Huajian
 */

package mysql

import (
	_ "github.com/go-sql-driver/mysql"
)

const (
	sqlAuthorTableCreate = iota
)

type (
	Author struct {
		name     string
		uri      string
	}

	authorServiceProvider struct {
		store *store
	}
)

var (
	sqlAuthor = []string {
		`CREATE TABLE IF NOT EXISTS feed (
			name     VARCHAR(512) NOT NULL,
			uri  VARCHAR(512) NOT NULL ,
		 )ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8_bin;`,
	}
)

//func InitAuthorMysql() {
//	db, err := sql.Open("mysql",
//		"root:123456@tcp(127.0.0.1:3306)/rss") //func Open(driverName, dataSourceName string) (*DB, error)
//
//	_, err = db.Exec(sqlEntry[sqlAuthorTableCreate])
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//}

func (a *authorServiceProvider) CreateTable() error {
	_, err := a.store.db.Exec(sqlAuthor[sqlAuthorTableCreate])

	return err
}

func (a *authorServiceProvider) Create(name, uri string) error {
	return nil
}
