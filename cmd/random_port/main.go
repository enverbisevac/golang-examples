package main

import (
	"log"
	"net"
	"net/http"
)

func createListener() (net.Listener, func()) {
	// port 0 tells the system to find any available port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	return l, func() {
		_ = l.Close()
	}
}

func main() {
	l, close := createListener()
	defer close()

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healty"))
	}))

	log.Println("Serving at", l.Addr().(*net.TCPAddr).Port)
	http.Serve(l, nil)
}
