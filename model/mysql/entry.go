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
	`CREATE TABLE IF NOT EXISTS feed (
		id   VARCHAR(512) NOT NULL,
		title     VARCHAR(512) NOT NULL,
		link      VARCHAR(512) NOT NULL,
		published TIMESTAMP(6),
		updated   TIMESTAMP(6)
	 )ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8mb4 COLLATE=utf8_bin;`,
	}
)

//func InitEntryMysql() {
//db, err := sql.Open("mysql",
//"root:123456@tcp(127.0.0.1:3306)/rss") //func Open(driverName, dataSourceName string) (*DB, error)
//
//_, err = db.Exec(sqlEntry[sqlEntryTableCreate])
//
//if err != nil {
//log.Fatal(err)
//}
//defer db.Close()
//}
func (f *entryServiceProvider) CreateTable() error {
	_, err := f.store.db.Exec(sqlFeed[sqlFeedTableCreate])

	return err
}

func (f *entryServiceProvider) Create(id uint32, title, subtitle string, up string) error {
	return nil
}
