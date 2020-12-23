package db

import (
	"database/sql"
	"fmt"
)

//DB will be used for db business
var DB *sql.DB

//Config is to start connection with db
func Config() {
	host := "127.0.0.1"
	port := "5432"
	user := "yourname"
	password := "yourpassword"
	dbname := "yourdbname"

	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database connection is succesfully")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
