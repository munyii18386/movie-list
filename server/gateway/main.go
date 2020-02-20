package main

import (
	"time"
	"encoding/json"
	"io/ioutil"
	// "github.com/gorilla/mux"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"


	// h "movie-list/server/gateway/handlers"
	
)

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}
// HelloHandler is a
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// SignupHandler .........
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am here")
	var u NewUser
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Email)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("sign up!"))
}

//SetHeader something
type SetHeader struct {
	Handler http.Handler
}

//ServeHTTP something
func (sh *SetHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSFR-Token, Authorization")
	// w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.Header().Set("Access-Control-Max-Age", "600")
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	start := time.Now()
	sh.Handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.String(), time.Since(start))
}

//NewSetHeader constructs a new Logger middleware handler
func NewSetHeader(handlerToWrap http.Handler) *SetHeader {
	return &SetHeader{handlerToWrap}
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
	

	r := http.NewServeMux()
	r.HandleFunc("/hello", HelloHandler)
	r.HandleFunc("/api/signup", SignupHandler)
	r.Handle("/", reactProxy)
	wm := NewSetHeader(r)
	fmt.Printf("listening on %s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wm))

}
