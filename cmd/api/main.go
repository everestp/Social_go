package main

import (
	"log"

	"github.com/everestp/Social_go/internal/store"
)


func main(){
	cfg := config{
		addr: ":8080",
	}
	store := store.NewPostgressStorage(nil)
	app := &application{
		config: cfg,
		store: store,

	}
	 
	mux :=app.mount()


	log.Fatal(app.run(mux))
}