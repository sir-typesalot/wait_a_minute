package util

import (
	"fmt"
    "log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB_Conn struct {
	db *sqlx.DB
	CONN_OPEN bool
}

func connectToDB(DB *DB_Conn) (outputDB *sqlx.DB, outputError error) {

	var err error
	// Create connection string
	connection_string := connectionString()
	DB.db, err = sqlx.Open("mysql", connection_string)
	// handle connection error
	if err != nil {
		log.Fatal(err)
	}
	DB.CONN_OPEN = false
	outputDB = DB.db
	outputError = err
	// return connection
    return
}

func GetConnection() *sqlx.DB {

	dbCon, errMessage := connectToDB(&DB_Conn{})
	if errMessage != nil {
		log.Fatal("Error Connecting to DB")
	}
	return dbCon
}

func connectionString() string {
	dbPass := VaultGet("DB_PASS")
    dbUser := VaultGet("DB_USER")
	dbHost := VaultGet("DB_HOST")
	connection_string := fmt.Sprintf("%s:%s@tcp(%s:3306)/wam", dbUser, dbPass, dbHost)
	return connection_string
}
