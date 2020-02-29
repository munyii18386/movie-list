package database

import(
	"database/sql"
	"errors"
	"fmt"
	"movie-list/server/gateway/handlers"
)

type instance struct{
	CONN *sql.DB
}

// CreateInstance ............
func CreateInstance(conn *sql.DB) handlers.DatabaseStore {
	return &instance{
		CONN: conn,
	}
}

func (obj *instance) Insert(person *handlers.User) (*handlers.User, error){
	q := "insert into User(Email, Passhash, FirstName, LastName) values (?,?,?,?)"
	result, err := obj.CONN.Exec(q, person.Email, string(person.PassHash), person.FirstName, person.LastName)
	if (err != nil){
		return nil, err
	}
	id, _ := result.LastInsertId()
	user, err := obj.LocateID(id)
	return user, err
}

func (obj *instance) LocateID(id int64) (*handlers.User, error) {
	user := handlers.User{}

	if id < 0 {
		return nil, errors.New("invalid primary key")
	}

	rows := obj.CONN.QueryRow("SELECT * FROM User WHERE id=?", id)

	if err := rows.Scan(&user.ID, &user.Email, &user.PassHash, &user.FirstName, &user.LastName); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s", "User not found")
		}
		return nil, fmt.Errorf("error scanning row: %v", err)
	}
	return &user, nil
}
