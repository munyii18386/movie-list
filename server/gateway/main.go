package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
	h "movie-list/server/gateway/handlers"
	"strings"
	"sync"
	"sync/atomic"

)

// HelloHandler is a
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

//  getURLS takes in strings of URL, splits and parses each string. It returns a list of URLS
func getURLS(targets string) []*url.URL{
	splitAddrs := strings.Split(targets, ",")
	myURLs := make([]*url.URL, len(splitAddrs))
	for i, c := range splitAddrs{
		URL, err := url.Parse(c)
		if err != nil { log.Fatal(fmt.Printf("Failed to parse URL %v", err))}
		fmt.Printf("parsed URL is: %s", URL)
		myURLs[i] = URL
	}
	return myURLs
}

// NewServiceProxy func ....
func NewServiceProxy(addrs []*url.URL) *httputil.ReverseProxy{
	var nextIndex int32
	nextIndex = 0
	mutex := sync.Mutex{}
	return &httputil.ReverseProxy{
		Director: func(r *http.Request){
			mutex.Lock()
			defer mutex.Unlock()
			target := addrs[nextIndex%int32(len(addrs))]
			atomic.AddInt32(&nextIndex, 1)
			r.Host = target.Host
			r.URL.Host = target.Host
			r.URL.Scheme = target.Scheme
		},
	}

}
// SetHeader something 
type SetHeader struct{
	Handler http.Handler
}

// ServeHTTP something
func (sh *SetHeader) ServeHTTP(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.Header().Set("Access-Control-Max-Age", "600")
	if r.Method == "OPTIONS"{
		w.WriteHeader(http.StatusOK) 
		return
	}
	start := time.Now()
	sh.Handler.ServeHTTP(w,r)
	log.Printf("%s %s %v", r.Method, r.URL.String(), time.Since(start))
}

// NewSetHeader function is ..
func NewSetHeader(wrapThisHandler http.Handler) *SetHeader{
	return &SetHeader{wrapThisHandler}
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

	// reactURLS := getURLS(reactAddr)

	reactProxy := httputil.NewSingleHostReverseProxy(&url.URL{Scheme: "http", Host: reactAddr})
	// reactProxy := NewServiceProxy(reactURLS)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	mux.HandleFunc("/signup", h.SignUp)
	mux.Handle("/", reactProxy)

	wrappedMux := NewSetHeader(mux)

	fmt.Printf("listening on %s...\n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))

}
