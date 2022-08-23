package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type postgresRespository struct {
	uri string
	db  *sql.DB
}

func NewPostgresRespository(uri string) *postgresRespository {
	return &postgresRespository{
		uri: uri,
	}
}

func (r *postgresRespository) Connect() {
	db, err := sql.Open("postgres", r.uri)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	r.db = db

	log.Println("successfully connected!")
}

func (r *postgresRespository) Disconnect() {
	log.Println("closing postgres connection!")

	r.db.Close()

	log.Println("successfully closed!")
}

func (r *postgresRespository) GetDB() *sql.DB {
	return r.db
}
