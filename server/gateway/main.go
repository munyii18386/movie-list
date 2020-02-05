package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// HelloHandler is a
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	addr := os.Getenv("ADDR")
	reactAddr := os.Getenv("REACT")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")

	fmt.Printf("reactAddr: %s", reactAddr)

	reactProxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: reactAddr})

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	mux.Handle("/", reactProxy)
	fmt.Printf("listening on %s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))

}
