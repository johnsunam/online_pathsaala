package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type Database struct {
	Db  *sql.DB
	Log *zap.SugaredLogger
}

func ConnectDb() (*sql.DB, error) {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	userName := os.Getenv("USERNAME")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	psqlUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", userName, password, host, port, dbname)
	dbs, errsql := sql.Open("postgres", psqlUrl)
	return dbs, errsql
}
