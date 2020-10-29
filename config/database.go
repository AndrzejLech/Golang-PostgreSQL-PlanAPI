package config

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	psqlInfo := "postgres://postgres:gopher@postgres:5432?sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil{
	panic(err)
	}

	err = db.Ping()
		if err != nil{
			panic(err)
		}
	return db
}
