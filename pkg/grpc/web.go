package grpc

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func WebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/cron", func(w http.ResponseWriter, r *http.Request) {
		command := r.URL.Query().Get("command")
		ClientConnect("4040", command)
		ClientConnect("4041", command)
		fmt.Fprintf(w, "Job Scheduled")
	})
	
	log.Fatal(http.ListenAndServe(":8081", nil))
}
