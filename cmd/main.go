package main

import (
	"log"
	"github.com/joho/godotenv"
	"github.com/Magdi888/EcoWithGo/cmd/api"
	"github.com/Magdi888/EcoWithGo/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	godotenv.Load(".env")
	DatabaseConn := database.EnvDBConfig()
	db, err := database.NewMySQLConfig(mysql.Config{
		User:   DatabaseConn.Username,
		Passwd: DatabaseConn.Password,
		Addr:   DatabaseConn.DBAddress,
		DBName: DatabaseConn.Database,
		Net:    "tcp",
		AllowNativePasswords:  true, //TODO remove this for production!
		ParseTime:       true,
	})

	if err != nil  {
		log.Fatal(err)
	}

	database.InitDB(db)

	server := api.NewApiServer(":8080", db)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

