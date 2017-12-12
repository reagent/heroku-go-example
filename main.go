package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	handler := http.NewServeMux()

	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		body := `{"status":"ok"}` + "\n"

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		io.WriteString(w, body)
		log.Printf("\"GET / %s\" %d %d\n", r.Proto, status, len(body))
	})

	log.Printf("Starting server on port %s\n", port)

	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
