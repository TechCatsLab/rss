/*
 *  Revision History:
 *      Initial: 2018/08/14    Wang Huajian
 */

package mysql

import (
	"database/sql"
)

// StoreServiceProvider encapsulated database
type StoreServiceProvider struct {
	store *store
}

// operate database
type store struct {
	db *sql.DB

	feedService *feedServiceProvider
	entryService *entryServiceProvider
}

var (
	StoreService *StoreServiceProvider
)

// Initial store or Instantiated store
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

	StoreService = &StoreServiceProvider{
		store: s,
	}
}

// Mount a method(FeedServiceProvider) for StoreServiceProvider struct
func (s *StoreServiceProvider) FeedServiceProvider() *feedServiceProvider {
	return s.store.feedService
}

// same as above
func (s *StoreServiceProvider) EntryServiceProvider() *entryServiceProvider {
	return s.store.entryService
}
