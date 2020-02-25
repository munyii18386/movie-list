package models

//	"movie-list/server/gateway/handlers"

//User represents a user account in the database
type User struct {
	ID        int64  `json:"id"`
	Email     string `json:"-"` //never JSON encoded/decoded
	PassHash  []byte `json:"-"` //never JSON encoded/decoded
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}



// Hub represents ......
type Hub interface {
	Insert(user *User) (*User, error)
}
