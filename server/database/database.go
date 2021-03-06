package database

import(
	"database/sql"
	"errors"
	"fmt"
	"movie-list/server/gateway/handlers"
	"log"
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

func (obj *instance) InsertMovie(umi *handlers.UserMovieInfo) (*handlers.UserMovieInfo, error) {
	q := "insert into Movie (UserID,  MovieURL, MovieTitle, MovieOverview) values (?,?,?,?)"
	result, err := obj.CONN.Exec(q, umi.UserID, umi.URL, umi.Title, umi.Overview)
	if (err != nil){
		return nil, err
	}
	id, _ := result.LastInsertId()
	m , err := obj.LocateMovieByID(id)
	return m, err
}

func (obj *instance) GetAllMovies(id int64) ([]handlers.UserMovieInfo, error) {
	var list []handlers.UserMovieInfo
	m := handlers.UserMovieInfo{}

	q := "select ID, UserID, MovieURL, MovieTitle, MovieOverview from Movie where UserID = ?"
	rows, err := obj.CONN.Query(q, id)
	if err != nil{
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&m.MovieID, &m.UserID,  &m.URL, &m.Title, &m.Overview ); err != nil {
			if err == sql.ErrNoRows {
				return nil, fmt.Errorf("%s", "Movie not found")
			}
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		fmt.Printf(" movie item scanned is : %+v\n", m)
		list = append(list, m)
		
	}

	for i, val := range list {
		fmt.Printf("%d\n movie item in db is: %+v\n", i, val)
	}
	return list, err
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

func (obj *instance) LocateMovieByID(id int64) (*handlers.UserMovieInfo, error) {
	movie := handlers.UserMovieInfo{}
	if id < 0 {
		return nil, errors.New("invalid primary key")
	}
	rows := obj.CONN.QueryRow("SELECT * FROM Movie WHERE id=?", id)
	if err := rows.Scan(&movie.MovieID, &movie.UserID,  &movie.URL, &movie.Title, &movie.Overview ); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%s", "Movie not found")
		}
		return nil, fmt.Errorf("error scanning row: %v", err)
	}
	return &movie, nil
}

// Delete removes a User with a given id from the db
func (obj *instance) Delete(id int64) error {
	_, e := obj.CONN.Exec("DELETE FROM User WHERE ID=?", id)
	if e != nil {
		return e
	}
	return e
}
