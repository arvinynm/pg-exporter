package conn

import (
	"database/sql"
	"fmt"
)

const (
	host     = "192.168.220.129"
	port     = "5432"
	user     = "postgres"
	password = "123456"
	dbname   = "postgres"
)

func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
