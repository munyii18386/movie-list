package handlers

import(
	"fmt"
	"net/http"
	"time"
	"io/ioutil"
	"log"
	"strconv"
	"encoding/json"
)


// WallHandler is used to
func (ctx *Context) WallHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Wall route reached!")
	if r.Method == http.MethodPost {
		token := r.Header.Get("Authorization")
		fmt.Println("Auth: ", token)

		// Loop over header names
	// for name, values := range r.Header {
    // 	// Loop over all values for the name.
    // 	for _, value := range values {
    //    		 fmt.Println(name, value)
   	// 	 }
	// }
	user := &User{
		ID: 0,
		Email: "",
		PassHash: []byte(""),
		FirstName: "",
		LastName: "",
	}
	state := SessionState{SessionTime: time.Now(), User: user}
	result, err := GetState(token, ctx.SessionKey, ctx.RedisDatabase, state)
	HandleError(err)
	rr, err := json.Marshal(result)
	HandleError(err)
	ss := SessionState{}
	err = json.Unmarshal(rr, &ss)
	HandleError(err)
	fmt.Printf("session state retrived  is: %+v\n", result)
	// fmt.Printf("user  is: %+v\n", ss.User)

	if err == nil {
		userID := ss.User.ID
		fmt.Println("userID is: ", userID)
		var m MovieDetail
		// get body
		body, err := ioutil.ReadAll(r.Body)
		ExitTransaction(err)
		log.Println(string(body))
		err = json.Unmarshal(body, &m)
		ExitTransaction(err)
		// fmt.Printf("movie detail is: %+v\n", m)

		HandleError(err)
		userMovieInfo := &UserMovieInfo{
			UserID: ss.User.ID,
			URL: m.URL,
			Title: m.Title,
			Overview: m.Overview,
		}
		// fmt.Printf("user movie detail is: %+v\n", userMovieInfo)

		updatedResult, err := ctx.UserDatabase.InsertMovie(userMovieInfo)
		fmt.Printf("updated user movie detail is: %+v\n", updatedResult)

		send := &Info {
			MovieID: strconv.FormatInt(updatedResult.MovieID, 10),
			UserID: strconv.FormatInt(updatedResult.UserID, 10),
			MovieURL: updatedResult.URL,
			Title: updatedResult.Title,
			Overview: updatedResult.Overview,
			MovieAdded: true,
		}

		w.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(send)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
		
	} else{
		// send err back to browser
	}


	}

}