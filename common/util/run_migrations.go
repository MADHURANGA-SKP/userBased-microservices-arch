package util

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
)

func RunDbMigrations(migrationURL string, dbsoruce string){
	migration, err := migrate.New(migrationURL, dbsoruce)
	if err != nil {
		log.Fatal(err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run db migration up")
	}

	log.Printf("db migrated successfully")
}