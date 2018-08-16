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
	store *store // 为什么类型写 *store, 直接写 store   store  不行吗？
}

// 操作数据库
type store struct {
	db *sql.DB

	feedService *feedServiceProvider

	entryService *entryServiceProvider

	authorService *authorServiceProvider
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

	s.entryService = &entryServiceProvider{
		store: s,
	}
	if err := s.entryService.CreateTable(); err != nil {
		panic(err)
	}

	s.authorService = &authorServiceProvider{
		store: s,
	}
	if err := s.authorService.CreateTable(); err != nil {
		panic(err)
	}

	StoreService = &StoreServiceProvider{
		store: s,
	}
}

func (s *StoreServiceProvider) FeedServiceProvider() *feedServiceProvider {
	return s.store.feedService
}

func (s *StoreServiceProvider) EntryServiceProvider() *entryServiceProvider {
	return s.store.entryService
}

func (s *StoreServiceProvider) AuthorServiceProvider() *authorServiceProvider {
	return s.store.authorService
}
