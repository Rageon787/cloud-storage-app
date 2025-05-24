package database

import (
	"database/sql"
	"fmt"
)

type DB struct {
	Conn *sql.DB
}

func connectDB() *DB {

	conn, err := sql.Open("sqlite3", "./app.db") // try to connect to the app's database
	if err != nil {
		fmt.Print("Connection failed") // if it fails to connect, log an error (will implement logging later)
	}

	return &DB{
		Conn: conn,
	}
}
