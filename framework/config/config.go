package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var Database *sqlx.DB

// DatabaseConnection contain the settings of the database
type DatabaseConnection struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

// Init is implementation set environment variables
func Init() {

	setupDatabase()

}

func mustGetenv(key string) string {
	valorAmbiente := os.Getenv(key)
	if valorAmbiente == "" {
		log.Fatalf("%s variável de ambiente não localizado.", key)
	}
	return valorAmbiente
}
func mustGetenvInt(key string) int {
	valorAmbiente := os.Getenv(key)
	valorAmbienteInt, err := strconv.Atoi(valorAmbiente)
	if err != nil {
		log.Fatalf("%s variável de ambiente não localizado.", key)
	}
	return valorAmbienteInt

}

func setupDatabase() {

	dbConfig := DatabaseConnection{
		Host:     mustGetenv("PSQL_DB_HOST"),
		Port:     mustGetenvInt("PSQL_DB_PORT"),
		User:     mustGetenv("PSQL_DB_USER"),
		Password: mustGetenv("PSQL_DB_PASSWORD"),
		Dbname:   mustGetenv("PSQL_DB_NAME"),
	}
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)

	Database = newConnectionPSQL(psqlConn)

}

func newConnectionPSQL(psqlConn string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", psqlConn)
	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(30 * time.Minute)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
