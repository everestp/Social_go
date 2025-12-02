package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)


var (
	ErrNotFound = errors.New("Resource not found")
	QueryTimeoutDuration  = time.Second * 5
)



type Storage struct {
	Posts  interface {
		Create(context.Context  ,*Post) error
		GetByID(context.Context , int64) (*Post ,error)
		Delete(context.Context , int64) error 
		Update(context.Context , *Post) error 
	}
	User  interface {
		GetByID(context.Context , int64) (*User ,error)
		Create(context.Context , *User) error
	}
	Comments interface{
			Create(context.Context  ,*Comment) error
		GetByPostID(ctx context.Context ,postID int64 )([]Comment ,error)
	}

}


func NewPostgressStorage(db *sql.DB ) Storage {
	
	return  Storage{
         Posts: &PostStore{db},
		 User:&UserStore{db},
		 Comments: &CommentStore{db},

	}
}