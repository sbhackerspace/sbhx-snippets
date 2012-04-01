// Steve Phillips / elimisteve
// 2012.03.31

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func showStopper(where string, err error) {
	if err != nil {
		fmt.Printf("Error near %s: %s\n", where, err)
		return
	}
}

func main() {
	const DB_NAME = "./gotest.db"

	db, err := sql.Open("sqlite3", DB_NAME)
	showStopper("open", err)
	defer os.Remove(DB_NAME)

	// _, err = db.Exec("drop table foo")
	create := "create table foo (id integer, name string, awesomeness integer)"
	_, err = db.Exec(create)
	showStopper("create table foo", err)

	// Insert data
	_, err = db.Exec("insert into foo(id, name, awesomeness) values(1, ?, ?)",
		"steve", 8)
	_, err = db.Exec("insert into foo(id, name, awesomeness) values(2, ?, ?)",
		"aj", 9)
	showStopper("insert steve and aj into foo", err)

	// Find inserted values
	rows, err := db.Query("select name, awesomeness from foo")
	showStopper("select stuff from foo", err)
	defer rows.Close()
	for rows.Next() {
		var awesomeness int
		var name string
		rows.Scan(&name, &awesomeness)
		fmt.Printf("%s's current awesomeness level: \t %d\n", name, awesomeness)
	}
}
