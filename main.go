package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/namnguyen191/github-api-server/api"
)

const defaultMainPort = "8000"

type appConfig struct {
	mainPort string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load env file")
	}

	app := &appConfig{}

	http.HandleFunc("/", api.Handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultMainPort
	}
	app.mainPort = port

	fmt.Println("Starting listening on port: ", app.mainPort)
	err = http.ListenAndServe(":"+app.mainPort, nil)
	if err != nil {
		panic(err)
	}
}
