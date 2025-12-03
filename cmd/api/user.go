package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/everestp/Social_go/internal/store"
	"github.com/go-chi/chi/v5"
)


type userKey string
const userCtx userKey ="user"

func (app *application) getUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse userID from URL
	userID, err := strconv.ParseInt(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	// Get context
	ctx := r.Context()

	// Fetch user from store
	user, err := app.store.User.GetByID(ctx, userID)
	if err != nil {
		switch err {
		case store.ErrNotFound:
			app.badRequestResponse(w, r, err)
			return
		default:
			app.internalServerError(w, r, err)
			return
		}
	}

	// Return JSON response
	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}
  type FollowUser struct {
	UserID int64 `json:"user_id"`

  }

func (app *application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	followerUser := getUserFromContext(r)
// TODO Revert Back to auth UserID from ctx
var payload FollowUser
if err := readJSON(w, r, &payload); err !=nil{
	app.badRequestResponse(w, r , err )
}
  if err := app.store.Followers.Follow(r.Context(), followerUser.ID, payload.UserID); err != nil{
	app.internalServerError(w, r , err )
  }
}

func (app *application) unfollowUserHandler(w http.ResponseWriter, r *http.Request){
unfollowedUser := getUserFromContext(r)

// TODO Revert back to auth UserID from ctx
 var payload FollowUser
 if err := readJSON(w, r , &payload); err != nil{
	app.badRequestResponse(w, r , err )
	return
 }

 ctx := r.Context()
if  err := app.store.Followers.Unfollow(ctx ,unfollowedUser.ID ,payload.UserID); err != nil{
	app.internalServerError(w, r, err)
		return
}
if err := app.jsonResponse(w, http.StatusNoContent, nil); err != nil {
		app.internalServerError(w, r, err)
		return
	}

} 


func (app *application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "userID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.badRequestResponse(w, r, fmt.Errorf("invalid post ID"))
			return
		}

		user, err := app.store.User.GetByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				app.notFoundResponse(w, r, err)
			} else {
				app.internalServerError(w, r, err)
				

			}
			
			return
		}

		ctx := context.WithValue(r.Context(), userCtx, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}




func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}
