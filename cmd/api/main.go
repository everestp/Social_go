package main

import (
	"log"

	"github.com/everestp/Social_go/internal/db"
	"github.com/everestp/Social_go/internal/store"
	
)
const version = "0.0.1"

func main(){
	cfg := config{
		addr: ":8001",
		db: dbConfig{
			Addr: "postgresql://neondb_owner:npg_k0WD6GpeTSui@ep-dawn-leaf-ad26dewr-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require",
			MaxOpenConns: 30,
			MaxIdleConns: 30,
			MaxIdleTime: "15m",
		},
		env: "development",
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