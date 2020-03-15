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

// GetByEmail returns a user with a given email
func (obj *instance) GetByEmail(email string) (*handlers.User, error) {
	rows, e := obj.CONN.Query("SELECT * FROM User WHERE Email=?", email)
	if e != nil {
		return nil, errors.New("Failed to retrieve matching rows")
	}
	defer rows.Close()
	c := handlers.User{}
	for rows.Next() {
		if err := rows.Scan(&c.ID, &c.Email, &c.PassHash, 
			&c.FirstName, &c.LastName); err != nil {
			fmt.Printf("error scanning row: %v\n", err)
		}
	}
	return &c, e
}


// Delete removes a User with a given id from the db
func (obj *instance) Delete(id int64) error {
	_, e := obj.CONN.Exec("DELETE FROM User WHERE ID=?", id)
	if e != nil {
		return e
	}
	return e
}
