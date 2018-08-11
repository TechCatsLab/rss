/*
 *  Revision History:
 *      Initial: 2018/08/10    Wang Huajian
 */

package database

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() {
	db, err := sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/rss") //func Open(driverName, dataSourceName string) (*DB, error)

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS `feed` (`feed_title` VARCHAR (512) NOT NULL, `feed_link` VARCHAR (512) NOT NULL, `feed_subtitle` VARCHAR (512) NOT NULL);" +
			"`feed_updated` VARCHAR(512) NOT NULL")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `entry` (`entry_id` INT UNSIGNED AUTO_INCREMENT, `entry_title` VARCHAR (512) NOT NULL," +
		"`entry_published` VARCHAR (512) NOT NULL, `entry_updated` VARCHAR (512) NOT NULL, `entry_summary` VARCHAR (4096) NOT NULL")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `author` (`author_name` VARCHAR (40) NOT NULL, `author_uri` VARCHAR (100) NOT NULL)ENGINE=InnoDB DEFAULT CHARSET=utf8;")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
