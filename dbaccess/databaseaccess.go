package dbaccess

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	_ "github.com/proullon/ramsql/driver"
)

var dbName = "userDb"

func InitDb() *sql.DB {

	db, err := sql.Open("ramsql", dbName)
	if err != nil {
		log.Error("sql.Open : Error : %s\n", err)
	}

	createSchema(db)

	return db
}

func createSchema(db *sql.DB)  {
	var userTable = "CREATE TABLE IF NOT EXISTS USERS (" +
		"idStudent INT NOT NULL PRIMARY KEY," +
		"name TEXT" +
		"surname TEXT)"
	log.Printf("Creating schema")

	statement, err := db.Prepare(userTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
}
