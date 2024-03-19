package config

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)


var db *sql.DB

func DatabaseInit() {
	var err error

	// db, err = sql.Open("mysql", "user=root dbname=goapi")
	db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/goapi")


	if err != nil {
		log.Fatal(err)
	}

	// Create Table cars if not exists
	createCarsTable()
}

// func createCarsTable() {
// 	_, err := db.Exec("CREATE TABLE IF NOT EXISTS cars(id serial,manufacturer varchar(20), design varchar(20), style varchar(20), doors int, created_at timestamp default NULL, updated_at timestamp default NULL, constraint pk primary key(id))")

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
func createCarsTable() {
    _, err := db.Exec("CREATE TABLE IF NOT EXISTS cars(id INT AUTO_INCREMENT, manufacturer VARCHAR(20), design VARCHAR(20), style VARCHAR(20), doors INT, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, PRIMARY KEY(id))")

    if err != nil {
        log.Fatal(err)
    }
}


// Getter for db var
func Db() *sql.DB {
	return db
}