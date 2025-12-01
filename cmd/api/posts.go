package main

import (
	"net/http"

	"github.com/everestp/Social_go/internal/store"
)

 type CreatePostPayload struct {
	 Title     string   `json:"title"`
Content   string   `json:"content"`
Tags      []string `json:"tag"`
 }

func (app *application) createPostHandler(w http.ResponseWriter , r *http.Request){
	var payload CreatePostPayload
	if err := readJSON(w, r , &payload); err != nil{
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return 
	}
	UserID :=1
	post := &store.Post{
		Title: payload.Title,
		Content: payload.Content,
		Tags: payload.Tags,
		//TODO Change after auth
		UserID: int64(UserID),
	}
ctx := r.Context()

	 if err := app.store.Posts.Create(ctx,post); err !=nil{
	writeJSONError(w, http.StatusBadRequest, err.Error())
		return 
	 }
  if err := writeJSON(w, http.StatusCreated, post); err != nil {
	writeJSONError(w, http.StatusBadRequest, err.Error())
		return 
  }
}