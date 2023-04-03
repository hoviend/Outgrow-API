package db

import (
	"fmt"
	"log"
	"outgrow/config"

	ent "outgrow/ent"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

func NewDBClient(config config.Config) *ent.Client {
	client, err := ent.Open(
		dialect.Postgres,
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", config.DBHost, config.DBPort, config.DBUser, config.DBName, config.DBPassword),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	return client
}
