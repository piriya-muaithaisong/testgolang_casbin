package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/piriya-muaithaisong/testgolang_casbin/api"
	db "github.com/piriya-muaithaisong/testgolang_casbin/db/sqlc"
	"github.com/piriya-muaithaisong/testgolang_casbin/utils"
)

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")
	}

	// if config.Environment == "development" {
	// 	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// }

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}

	// runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)

	// redisOpt := asynq.RedisClientOpt{
	// 	Addr: config.RedisAddress,
	// }

	// taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	runGinServer(config, store)
}

func runGinServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ")
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ")
	}
}
