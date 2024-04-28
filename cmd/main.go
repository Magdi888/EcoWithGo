package main

import (
	"log"
	"github.com/magdi888/cmd/api"
)

func main() {

	server := api.NewApiServer(":8080", nil)

	if err := server.Start(); err != nil {
		log.Fatal(err) 
		} 
}
