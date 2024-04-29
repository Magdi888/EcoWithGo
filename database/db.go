package database

import (
	"database/sql"
	"fmt"
	"os"
	"log"
	"github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	DBAddress string
}

func EnvDBConfig() *Config {
	return  &Config{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Username: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("DATABASE"),
		DBAddress: fmt.Sprintf("%s:%s",os.Getenv("MYSQL_HOST"),os.Getenv("MYSQL_PORT")),
	}
}
func  NewMySQLConfig( cfg mysql.Config) (db *sql.DB, err error){
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}	
	return db, nil
}

func InitDB(db *sql.DB)  {
	err := db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to database: %s\n", err)
	} else {
		log.Println("Connected to the Database!")
	}
	
}