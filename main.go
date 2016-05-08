package main

import (
	"log"
	"net/http"
	"os"
)

type logWrapper struct {
	handler http.Handler
}

func (l logWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] - %s\n", r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

func main() {
	dir := os.Getenv("DIR")
	if dir == "" {
		dir = os.ExpandEnv("${HOME}/Public")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", logWrapper{fs})
	log.Printf("Serving on http://0.0.0.0:%s/\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
