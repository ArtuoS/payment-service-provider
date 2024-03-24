package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Context struct {
	DB *sqlx.DB
}

func NewContext() *Context {
	db, err := sqlx.Connect("postgres", "user=pguser dbname=psp sslmode=disable password=pgadmin host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	return &Context{
		DB: db,
	}
}
