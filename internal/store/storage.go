package store

import (
	"context"
	"database/sql"
	"errors"
)


var (
	ErrNotFound = errors.New("Resource not found")
)



type Storage struct {
	Posts  interface {
		Create(context.Context  ,*Post) error
		GetByID(context.Context , int64) (*Post ,error)
		Delete(context.Context , int64) error 
		Update(context.Context , *Post) error 
	}
	User  interface {
		Create(context.Context , *User) error
	}
	Comments interface{
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