package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "postgres://postgres:password@localhost:5432/scoreDb?sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}

	db.SetMaxOpenConns(10)

	tx := db.MustBegin()
	db.MustExec(Schema)
	tx.Commit()

	fmt.Println("connection to database was successfull")
	return db
}
