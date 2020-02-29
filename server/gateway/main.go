package main

import (
	"github.com/go-redis/redis"
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
	redisAddr := os.Getenv("REDISADDR")
	sessionKey := os.Getenv("SESSIONKEY")
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
	// create an instance
	conn := database.CreateInstance(db)

	//redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	redisInstance := execute.NewRedisInstance(rdb, 120000000000)

	ctx := execute.Context{UserDatabase: conn, SessionKey: sessionKey, RedisDatabase: redisInstance}
	
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
