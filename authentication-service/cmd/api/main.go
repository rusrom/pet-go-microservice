package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
)

const (
	webPort            = "8091"
	connectionAttempts = 10
	sleepInterval      = 1 * time.Second
)

type Config struct {
	Db     *sql.DB
	Models data.Models
}

func main() {
	log.Println("authentication service is starting...")

	conn, err := connectToDB()
	if err != nil {
		panic(err.Error())
	}

	app := Config{
		Db:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func connectToDB() (*sql.DB, error) {
	dsn := os.Getenv("DSN")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	for attempt := 0; attempt < connectionAttempts; attempt++ {
		err = db.Ping()
		if err == nil {
			log.Println("Connected to database")
			return db, nil
		}

		log.Printf("Connection attempt %d failed. Retrying in %s...\n", attempt+1, sleepInterval)
		time.Sleep(sleepInterval)
	}

	return nil, fmt.Errorf("could not establish database connection after %d attempts", connectionAttempts)
}
