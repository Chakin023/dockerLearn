package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", handlePage)

	// Get the port from the environment variable, or use a default port if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if PORT is not set
	}

	addr := ":" + port
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
    <p> Hi Docker, I pushed a new version! </p>
</body>
</html>`
	w.Write([]byte(page))
}
