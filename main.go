package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

type app struct{}

func (a *app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := http.StatusOK
	body := `{"status":"ok"}` + "\n"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	io.WriteString(w, body)
	log.Printf("\"%s %s %s\" %d %d\n", r.Method, r.URL.Path, r.Proto, status, len(body))
}

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	log.Printf("Starting server on port %s\n", port)

	if err := http.ListenAndServe(":"+port, &app{}); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
