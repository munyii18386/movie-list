package handlers

import (
	"fmt"
	"log"

	// "os"
	"encoding/json"
	"errors"
	"io/ioutil"
	"movie-list/server/gateway/models"
	"net/http"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

// //User represents a user account in the database
// type User struct {
// 	ID        int64  `json:"id"`
// 	Email     string `json:"-"` //never JSON encoded/decoded
// 	PassHash  []byte `json:"-"` //never JSON encoded/decoded
// 	FirstName string `json:"firstName"`
// 	LastName  string `json:"lastName"`
// }

// ConfUser ................
type ConfUser struct {
	ID     string
	Name   string
	Status string
}

// Context .................................
type Context struct {
	Users models.Hub
}

// SignUpHandler handles
func (ctx *Context) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var u NewUser
	var resp ConfUser

	fmt.Println("I am here")
	if r.Method == http.MethodPost {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))

		err = json.Unmarshal(body, &u)

		if err != nil {
			panic(err)
		}

		fmt.Println(u)

		// check if email is valid
		addr, err := mail.ParseAddress(u.FirstName + " " + u.LastName + "<" + u.Email + ">")
		if err != nil {
			err = errors.New("Invalid Email Address")
		}

		// check that passwords match
		if u.Password != u.PasswordConf {
			err = errors.New("Password and PasswordConf do not match")
		}
		// check password length
		if len(u.Password) < 6 {
			err = errors.New("Password is less than 6 characters")
		}
		// hash password
		passhash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 13)
		if err != nil {
			fmt.Printf("error generating bcrypt hash: %v\n", err)
		}
		fmt.Println("this is the password hash: ", string(passhash))

		user := &models.User{
			Email:     addr.Address,
			PassHash:  passhash,
			FirstName: u.FirstName,
			LastName:  u.LastName,
		}
		fmt.Printf("%+v\n", user)

		_, err = ctx.Users.Insert(user)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		resp.ID = "1"
		resp.Name = "Lilia"
		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

	}

}
