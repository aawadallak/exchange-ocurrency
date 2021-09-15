package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

type Database struct {
	databaseUser     string
	databasePassword string
	databaseSSL      string
	databaseName     string
	databaseHost     string
	databasePort     string
}

func createDatabase() Database {
	return Database{
		databaseUser:     os.Getenv("DB_USER"),
		databasePassword: os.Getenv("DB_PASS"),
		databaseSSL:      os.Getenv("DB_SSL"),
		databaseName:     os.Getenv("DB_NAME"),
		databaseHost:     os.Getenv("DB_HOST"),
		databasePort:     os.Getenv("DB_PORT"),
	}

}

var database *pgx.Conn

func Init() {

	_db := createDatabase()

	str := fmt.Sprintf("postgres://%s:%s@%s:5432/%s", _db.databaseUser, _db.databasePassword, _db.databaseHost, _db.databaseName)

	db, err := pgx.Connect(context.Background(), str)

	if err != nil {
		log.Fatalln(err)
	}

	database = db

}

func RunMigrations() error {

	_db := createDatabase()

	dsn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", _db.databaseUser, _db.databasePassword, _db.databaseHost, _db.databasePort, _db.databaseName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	_, err = postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalln(err)
		return err
	}

	m, err := migrate.New(
		"file://database/migrations/",
		dsn)

	if err != nil {
		log.Fatalln(err)
		return err
	}

	if err := m.Up(); m.Up() != nil {
		return err
	}

	return nil
}

func CloseConn(ctx context.Context) {
	database.Close(ctx)
}

func GetDatabase() *pgx.Conn {
	return database
}
