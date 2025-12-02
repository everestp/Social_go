package main

import (
	"log"

	"github.com/everestp/Social_go/internal/db"
	"github.com/everestp/Social_go/internal/store"
)

func main(){
addr:= "postgresql://neondb_owner:npg_k0WD6GpeTSui@ep-dawn-leaf-ad26dewr-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require";
	conn ,err := db.New(addr, 3, 3, "15m")
	if err !=nil {
		log.Fatal(err)
	}
	defer conn.Close()
	store :=store.NewPostgressStorage(conn) 
	db.Seed(store)
}
