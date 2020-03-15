package handlers

import (
	"time"
	"fmt"
	"log"

	// "os"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)






// SignUpHandler handles
func (ctx *Context) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var u NewUser
	
	fmt.Println("I am here")
	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)
		ExitTransaction(err)
		log.Println(string(body))

		err = json.Unmarshal(body, &u)
		ExitTransaction(err)

		fmt.Println(u)

		//check first name length
		if ((len(u.FirstName) < 1) || (u.FirstName == "")) {
			err = errors.New("First Name is less than 6 characters")
			HandleError(errors.New("First Name is less than 6 characters"))
		}
	
		// check last name length
		if ((len(u.LastName) < 1) || (u.LastName == "")) {
			err = errors.New("Last Name is less than 6 characters")
			HandleError(errors.New("Last Name is less than 6 characters"))
		}

		// check email length
		if len(u.Email) < 1 {
			err = errors.New("Email is less than 6 characters")
			HandleError(errors.New("Email is less than 6 characters"))
		}

		// check if email is valid
		addr, err := mail.ParseAddress(u.FirstName + " " + u.LastName + "<" + u.Email + ">")
		HandleError(err)

		// check that passwords match
		if u.Password != u.PasswordConf {
			err = errors.New("Password and PasswordConf do not match")
			HandleError(err)
		}
		// check password length
		if len(u.Password) < 6 {
			err = errors.New("Password is less than 6 characters")
			HandleError(err)
		}

		if (err == nil){
			// hash password
			passhash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 13)
			HandleError(err)
			fmt.Println("this is the password hash: ", string(passhash))
			user := &User{
				Email:     addr.Address,
				PassHash:  passhash,
				FirstName: u.FirstName,
				LastName:  u.LastName,
			}
			fmt.Printf("%+v\n", user)
			person, err := ctx.UserDatabase.Insert(user)
			HandleError(err)
	
			
			state := SessionState{SessionTime: time.Now(), User: person}
			fmt.Printf("STATE: %+v\n", state)
			sid, err := StartSession(ctx.SessionKey, ctx.RedisDatabase, state, w )
			fmt.Println("SID: ", sid)
			

			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
				return
			}

			profile := &Profile{
				ID: person.ID, 
				Email: person.Email,
				FirstName: person.FirstName,
				Status: "true",
			}
			fmt.Printf("the user info from the database: %+v\n", profile)

			w.WriteHeader(http.StatusCreated)
			w.Header().Add("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(profile)
			if err != nil {
				http.Error(w, "Bad request", http.StatusBadRequest)
			}
		}
		
	}

}
