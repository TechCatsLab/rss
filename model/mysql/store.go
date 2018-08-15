/*
 *  Revision History:
 *      Initial: 2018/08/14    Wang Huajian
 */

package mysql

import (
	"database/sql"
)

// StoreServiceProvider 作用：封装数据库
type StoreServiceProvider struct {
	store *store
}

// 操作数据库
type store struct {
	db *sql.DB

	feedService *feedServiceProvider
}

var (
	StoreService *StoreServiceProvider
)

// 初始化 store 或实例化 store
func InitStoreService(db *sql.DB) {
	s := &store{
		db: db,
	}

	s.feedService = &feedServiceProvider{
		store: s,
	}

	if err := s.feedService.CreateTable(); err != nil {
		panic(err)
	}

	StoreService = &StoreServiceProvider{
		store: s,
	}
}

func (s *StoreServiceProvider) FeedServiceProvider() *feedServiceProvider {
	return s.store.feedService
}
