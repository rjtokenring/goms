package dbaccess

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	_ "github.com/proullon/ramsql/driver"
	"strconv"
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
		"userid INT NOT NULL PRIMARY KEY," +
		"name TEXT," +
		"surname TEXT)"
	log.Printf("Creating schema")

	statement, err := db.Prepare(userTable)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec() // Execute SQL Statements
	if err != nil {
		log.Fatal(err.Error())
	}
	err = statement.Close()
	if err != nil {
		log.Fatal(err.Error())
	}

	var add = "INSERT INTO USERS (userid, name, surname) values (1,'Marco','Colombo')"
	statement, err = db.Prepare(add)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec() // Execute SQL Statements
	if err != nil {
		log.Fatal(err.Error())
	}
}

func DeleteUserByID(db *sql.DB, key int64)  {
	stm, err := db.Prepare("delete from USERS where userid=?")
	if err != nil {
		log.Error("Error deleting user "+strconv.FormatInt(key, 10), err)
		return
	}
	_, err = stm.Exec(key)
	if err != nil {
		log.Error("Error deleting user "+strconv.FormatInt(key, 10), err)
		return
	}
	stm.Close()
}

func GetUserByID(db *sql.DB, key int64) (idUsr int64, nm string, surnm string) {
	stm, err := db.Prepare("select userid, name, surname from USERS where userid=?")
	if err != nil {
		log.Error("Error selecting user "+strconv.FormatInt(key, 10), err)
	}
	rows, err := stm.Query(key)
	if err != nil {
		log.Error("Error selecting user "+strconv.FormatInt(key, 10), err)
	}
	defer stm.Close()

	var id int64
	var name string
	var surname string

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &surname)
		if err != nil {
			log.Fatal(err)
		}
	}

	return id, name, surname
}
