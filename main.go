package main

import (
	"log"
	"net/http"
	"os"

	"github.com/uilian/cartoon-redirect/cmd/api"
)

func main() {
	http.HandleFunc("/", api.RedirectHandler)
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "3000"
	}
	log.Print("Listening on: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
