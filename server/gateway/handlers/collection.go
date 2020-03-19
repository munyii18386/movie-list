package handlers

import (
	"time"
)


//NewUser represents a new user signing up for an account
type NewUser struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordConf string `json:"passwordConf"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

// Profile is a snapshot of the full user info
type Profile struct{
	ID        int64  `json:"id"`
	Email     string `json:"email"` //never JSON encoded/decoded
	FirstName string `json:"firstName"`
	Status 	  string `json:"status"`
}

//User represents a user account in the database
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"-"` //never JSON encoded/decoded
	PassHash  []byte `json:"-"` //never JSON encoded/decoded
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// LoginRequest represents the user email and password requesting authentication
type LoginRequest struct{
	Email        string `json:"email"`
	Password     string `json:"password"`
}

// Context .................................
type Context struct {
	UserDatabase DatabaseStore
	SessionKey string
	RedisDatabase RedisStore
}

// DatabaseStore interface is 
type DatabaseStore interface{
	LocateID(id int64) (*User, error)
	Insert(user *User) (*User, error)
	GetByEmail(email string) (*User, error)
	Delete(id int64) error
	InsertMovie(umi *UserMovieInfo) (*UserMovieInfo, error)
	LocateMovieByID(id int64) (*UserMovieInfo, error)
}

// SessionState stores session time and authenticated user
type SessionState struct {
	SessionTime time.Time   `json:"session_time"`
	User        *User `json:"user"`
}

type MovieDetail struct{
	ID  string  `json:"id"`
	URL string `json:"url"`
	Title string `json:"title"`
	Overview string `json:"overview"`
}

type UserMovieInfo struct{
	MovieID  int64  `json:"id"`
	UserID int64  `json:"id"`
	URL string `json:"url"`
	Title string `json:"title"`
	Overview string `json:"overview"`
}

type Info struct{
	MovieID  string `json:"movieid"`
	UserID string `json:"userid"`
	MovieURL string `json:"movie_url"`
	Title string `json:"title"`
	Overview string `json:"overview"`
}

//RedisStore represents a session data store.
type RedisStore interface {
	//Save saves the provided `sessionState` and associated SessionID
	Save(sid string, sessionState interface{}) error
	//Get populates `sessionState` with the data previously saved
	//for the given SessionID
	Get(sid string, sessionState interface{}) (interface{}, error)
	//Delete deletes all state data associated with the SessionID from the store.
	Delete(sid string) error
	// Find retrieves value by key
	Find(sid string)

}



