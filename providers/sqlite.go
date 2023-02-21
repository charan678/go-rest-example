package providers

import (
	"database/sql"
	"log"
)

type sqliteDB struct {
	host string
	port int
}

type ISqliteProvider interface {
	Conn() (*sql.DB, error)
}

func NewSqliteDBProvider(host string, port int) ISqliteProvider {
	return &sqliteDB{host: host, port: port}
}

func (this *sqliteDB) Conn() (*sql.DB, error) {
	//sql.Open("sqlite3", "file:locked.sqlite?cache=shared")
	//test.db
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return db, err
}
