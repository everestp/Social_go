package main

import (
	"log"

	"github.com/everestp/Social_go/internal/db"
	"github.com/everestp/Social_go/internal/store"
	
)


func main(){
	cfg := config{
		addr: ":8000",
		db: dbConfig{
			Addr: "postgres://postgres:rest@localhost:5432/postgres?sslmode=disable",
			MaxOpenConns: 30,
			MaxIdleConns: 30,
			MaxIdleTime: "15m",
		},
	}
	db ,err := db.New(
		cfg.db.Addr,
		cfg.db.MaxOpenConns,
		  cfg.db.MaxIdleConns,
		  cfg.db.MaxIdleTime,
		)
		if err != nil {
			log.Panic(err)
		}
		defer db.Close()
		log.Println("Database connction pool eatablished")
	store := store.NewPostgressStorage(db)
	app := &application{
		config: cfg,
		store: store,
		

	}
	 
	mux :=app.mount()


	log.Fatal(app.run(mux))
}