package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

type logWrapper struct {
	handler http.Handler
}

func (l logWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] - %s\n", r.Method, r.URL.Path)
	l.handler.ServeHTTP(w, r)
}

func printInterfaces(port int) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Print("Can't figure out your IP addresses: " + err.Error())
		return
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			if len(ipnet.Mask) == 4 {
				log.Printf("Listening on http://%s:%d/", ipnet.IP.String(), port)
			}
		}
	}
}

func main() {
	dir := os.Getenv("DIR")
	if dir == "" {
		dir = os.ExpandEnv("${HOME}/Public")
	}
	portStr := os.Getenv("PORT")
	port, _ := strconv.Atoi(portStr)
	if port == 0 {
		port = 8888
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", logWrapper{fs})
	log.Print("Serving directory " + dir)
	printInterfaces(port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}
