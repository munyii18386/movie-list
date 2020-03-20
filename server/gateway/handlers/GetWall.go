package handlers


import(
	"fmt"
	"net/http"
	"time"
	// "io/ioutil"
	// "log"
	"strconv"
	"encoding/json"
)


// GetWallHandler is used to
func (ctx *Context) GetWallHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Println("get wall info reached")
	if r.Method == http.MethodGet {
		token := r.Header.Get("Authorization")
		fmt.Println("Auth: ", token)

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
		fmt.Printf("user  is: %+v\n", ss.User)
		if err == nil {
			uid := ss.User.ID
			list, err := ctx.UserDatabase.GetAllMovies(uid)
			HandleError(err)
			// fmt.Printf(" movie list is: %+v\n", list)
			var bucket []*Info
			for _, val  := range list {
				// fmt.Printf("%d\n movie list is: %+v\n", i, val)
				send := &Info {
					MovieID: strconv.FormatInt(val.MovieID, 10),
					UserID: strconv.FormatInt(val.UserID, 10),
					MovieURL:val.URL,
					Title: val.Title,
					Overview: val.Overview,
				}
				// fmt.Printf(" movie list is: %+v\n", send)
				bucket = append(bucket, send)
		
			}

		

			w.Header().Add("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(bucket)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

		}







	}
}