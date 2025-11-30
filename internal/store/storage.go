package store

import (
	"context"
	"database/sql"
)




type Storage struct {
	Posts  interface {
		Create(context.Context) error
	}
	User  interface {
		Create(context.Context) error
	}

}


func NewPostgressStorage(db *sql.DB) Storage {
	
	return  Storage{
         Posts: &PostStore{db},
		 User:&UserStore{db},

	}
}