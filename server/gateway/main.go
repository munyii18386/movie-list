package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	execute "movie-list/server/gateway/handlers"
	"movie-list/server/database"
	_ "github.com/go-sql-driver/mysql"
	
)


// HelloHandler is a
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}




func main() {
	// ENVs
	addr := os.Getenv("ADDR")	
	if len(addr) == 0 {
		addr = ":443"
	}
	reactAddr := os.Getenv("REACT")
	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")
	dsn := os.Getenv("DSN")


	// client proxy
	reactProxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: reactAddr, Path: "/"})

	// database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("error opening database: %v\n", err)
        os.Exit(1)
	}
	defer db.Close()
	// test to ping database
	if err := db.Ping(); err != nil {
        fmt.Printf("error pinging database: %v\n", err)
    } else {
        fmt.Printf("successfully connected!\n")
	}
	
	hub := database.CreateInstance(db)

	ctx := execute.Context{Users: hub}
	
	// routes
	r := http.NewServeMux()
	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/api/SignUp", ctx.SignUpHandler)
	r.Handle("/", reactProxy)

	// cors wrapper
	wm := execute.NewSetHeader(r)
	
	fmt.Printf("listening on %s...\n", addr)
	fmt.Printf("reactAddr: %s", reactAddr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wm))

}
