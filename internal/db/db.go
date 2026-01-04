package db

import (
	"database/sql"
	"log"
	"study-stack/internal/db/migrations"
	"study-stack/internal/env"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDb() *sql.DB {
	db, err := sql.Open("pgx", env.Config.POSTGRES_CONNECTION)
	if err != nil {
		log.Fatalf("Falied to open db: %v\n", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Falied to ping db: %v\n", err)
	}

	if err := migrations.MigrateFS(db, migrations.FS, "."); err != nil {
		log.Fatalf("Failed To Run DB Migrations: %v", err)
	}
	return db
}
