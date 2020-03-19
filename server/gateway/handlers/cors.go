package handlers

import(
	"net/http"
	"time"
	"log"
)

//SetHeader something
type SetHeader struct {
	Handler http.Handler
}



//Origin,OPTIONS,, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers');

//ServeHTTP something
func (sh *SetHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "OPTIONS, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSFR-Token, Authorization")
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