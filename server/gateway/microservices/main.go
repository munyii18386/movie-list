package main

import(
	"os"
	"fmt"
	"log"
	"net/http"
	h "movie-list/server/gateway/handlers"
)

func main() {

	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/SignUp", h.SignUp)

	fmt.Printf("listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))

}