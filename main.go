package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yu-yk/recipes-restful/api"
)

func main() {
	host := os.Getenv("CLEARDB_DATABASE_HOST")
	user := os.Getenv("CLEARDB_DATABASE_USER")
	password := os.Getenv("CLEARDB_DATABASE_PASSWORD")
	dbName := os.Getenv("CLEARDB_DATABASE_DB")

	db, err := newSQLConn("mysql", user+":"+password+"@tcp("+host+":3306)/"+dbName+"?parseTime=true")
	// db, err := newSQLConn("mysql", "user:password@tcp(localhost:3306)/db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := api.NewServer(db)
	// s.Serve("localhost:8080")
	s.Serve(":" + os.Getenv("PORT"))
}

func newSQLConn(driverName, url string) (*sql.DB, error) {
	db, err := sql.Open(driverName, url)
	if err != nil {
		return nil, err
	}

	// check connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to %s!\n", driverName)

	return db, nil
}
