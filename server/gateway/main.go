package main

import(
	"net/http"
	"os"
	"fmt"
	"log"
)

// HelloHandler is a 
func HelloHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte ("Hello World!"))
}

func main()  {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath := os.Getenv("TLSKEY")
    tlsCertPath := os.Getenv("TLSCERT")

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	fmt.Printf("listening on %s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
	  
}