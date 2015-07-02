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
	fs := http.FileServer(http.Dir(os.ExpandEnv("${HOME}/Public")))
	http.Handle("/", logWrapper{fs})
	log.Println("Serving on http://0.0.0.0:8888/")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		panic(err)
	}
}
