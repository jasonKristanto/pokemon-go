package config

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
)

func PgSqlConnection() {
	pgConfig := Configuration.Database

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", pgConfig.Host, pgConfig.Port),
		User:     pgConfig.Username,
		Password: pgConfig.Password,
		Database: pgConfig.Database,
	})

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}
}
