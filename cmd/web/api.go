package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	builder "github.com/uilian/cartoon-redirect/internal/pkg/builder"
)

func main() {
	http.HandleFunc("/", RedirectHandler)
	port, present := os.LookupEnv("PORT")
	if !present {
		port = "3000"
	}
	log.Print("Listening on: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Split(r.URL.Path, "/")[1]
	period := r.URL.Query().Get("q")
	url := builder.BuildRedirectURL(path, period)
	log.Print("Redirecting to: ", url)
	http.Redirect(w, r, url, http.StatusSeeOther)
}
