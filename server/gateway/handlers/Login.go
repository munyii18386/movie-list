package handlers

import(
	"net/http"
	"log"
	"time"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"net/mail"
	"errors"
	"golang.org/x/crypto/bcrypt"
)




// LoginHandler is used to
func (ctx *Context) LoginHandler (w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Login route reached!")
	var userReq LoginRequest

	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		ExitTransaction(err)
		log.Println(string(body))

		err = json.Unmarshal(body, &userReq)
		ExitTransaction(err)

		// fmt.Printf("the user login req is: %+v\n", userReq)

		// check if email is valid
		addr, err := mail.ParseAddress(userReq.Email)
		HandleError(err)
		// fmt.Println(addr)

		// check password length
		if len(userReq.Password) < 6 {
			err = errors.New("Password is less than 6 characters")
			HandleError(err)
		}


		exitingUser, err := ctx.UserDatabase.GetByEmail(addr.Address)
		// fmt.Printf("existing user profile is: %+v\n", exitingUser)
		HandleError(err)
		
		exitingUserHash := []byte(exitingUser.PassHash)
		// fmt.Println("existing user hash:", exitingUserHash)


		 //compare a password against this hash
		 if err := bcrypt.CompareHashAndPassword(exitingUserHash, []byte(userReq.Password)); err != nil {
			// fmt.Println("password doesn't match stored hash!")
		} else {
			// fmt.Println("password is valid")
			state := SessionState{SessionTime: time.Now(), User: exitingUser}
			fmt.Printf("state is: %+v\n", state)

			sid, err := StartSession(ctx.SessionKey, ctx.RedisDatabase, state, w )
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}
			fmt.Println("SID: ", sid)
			profile := &Profile{
				ID: exitingUser.ID, 
				Email: exitingUser.Email,
				FirstName: exitingUser.FirstName,
				Status: "true",
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Add("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(profile)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}

		}







	}
}