package database

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // sql driver
)

type instance struct{
	CONN *sql.DB
}

