package main

import (
	"log"

	"github.com/everestp/Social_go/internal/db"
	"github.com/everestp/Social_go/internal/store"
	"go.uber.org/zap"
)
const version = "0.0.1"
//	@title			GopherSocial API
//	@description	API for GopherSocial, a social network for gohpers
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath					/v1
//
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description

func main(){
	cfg := config{
		addr: ":8002",
		db: dbConfig{
			Addr: "postgresql://neondb_owner:npg_k0WD6GpeTSui@ep-dawn-leaf-ad26dewr-pooler.c-2.us-east-1.aws.neon.tech/neondb?sslmode=require&channel_binding=require",
			MaxOpenConns: 30,
			MaxIdleConns: 30,
			MaxIdleTime: "15m",
		},
		env: "development",
	}

//Logger
logger := zap.Must(zap.NewProduction()).Sugar()
defer logger.Sync()
//Database



	db ,err := db.New(
		cfg.db.Addr,
		cfg.db.MaxOpenConns,
		  cfg.db.MaxIdleConns,
		  cfg.db.MaxIdleTime,
		)
		if err != nil {
			logger.Fatal(err)
		}
		defer db.Close()
		logger.Info("Database connction pool eatablished")
	store := store.NewPostgressStorage(db)
	app := &application{
		config: cfg,
		store: store,
		logger: logger,
		

	}
	 
	mux :=app.mount()


	log.Fatal(app.run(mux))
}