package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/everestp/Social_go/internal/store"
	"github.com/google/uuid"
)

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=3,max=72"`
}

type UserWithToken struct {
	*store.User
	Token string `json:"token"`
}


func (app *application) registerUserHandler(w http.ResponseWriter , r *http.Request) {
	var  payload RegisterUserPayload
	if err := readJSON(w, r,  &payload ); err != nil{
		app.badRequestResponse(w, r , err )
		return
	}
    if err := Validate.Struct(payload); err !=nil {
		app.badRequestResponse(w, r , err )
		return
	}
user := &store.User{
	Username: payload.Username,
		Email:    payload.Email,
}

	// hash the  user password
	 if err := user.Password.Set(payload.Password); err != nil {
		app.internalServerError(w, r, err )
		return
	 }
	 ctx :=r.Context()
	 plainToken := uuid.New().String()

	 //store
	 hash :=sha256.Sum256([]byte(plainToken))
	 hashToken := hex.EncodeToString(hash[:])

	// store the user
	
	  err := app.store.User.CreateAndInvite(ctx, user, hashToken,app.config.mail.exp)
	  //mail
	  if err != nil {
		switch err {
		case store.ErrDuplicateEmail:
			app.badRequestResponse(w, r , err )
		case store.ErrDuplicateUsername :
			app.badRequestResponse(w, r , err )
		default:
			app.internalServerError(w, r , err )
		}
		return
	  }
UserWithToken := UserWithToken{
	User: user,
	Token: plainToken,
}

	if err :=  app.jsonResponse(w, http.StatusCreated, UserWithToken); err != nil {
		app.internalServerError(w, r, err)
	}


}