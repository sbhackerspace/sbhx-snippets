// Steve Phillips / elimisteve
// 2013.09.22

package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"os"
	"time"
)

const (
	CreateTestTable     = `CREATE TABLE goprofiletest (
	Id        serial PRIMARY KEY NOT NULL,
	Name      varchar(100) NOT NULL,
	Age       integer NOT NULL,
	CreatedAt timestamp NOT NULL DEFAULT current_timestamp
);`
	InsertProfile       = "INSERT INTO goprofiletest (Id, Name, Age, CreatedAt) VALUES (DEFAULT, $1, $2, DEFAULT)"
	SelectProfileByName = "SELECT Id, Name, Age, CreatedAt FROM goprofiletest WHERE Name = $1"
	SelectProfileByID   = "SELECT Id, Name, Age, CreatedAt FROM goprofiletest WHERE Id = $1"
	DeleteProfileByID   = "DELETE FROM goprofiletest WHERE id = $1"
	DropTestTable       = "DROP TABLE goprofiletest"
)

var (
	// Run this: PGPASS="MyPostgresPassword" go run postgres-insert-find.go
	PostgresConnStr = "user=postgres dbname=postgres host=localhost sslmode=disable password=" + os.Getenv("PGPASS")
)

func main() {
	db, err := sql.Open("postgres", PostgresConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create "goprofiletest" table, then drop it when we're done
	_, err = db.Exec(CreateTestTable)
	maybePanic(err)
	defer db.Exec(DropTestTable)


	// Create new Profile for elimisteve and insert into DB
	elimi := Profile{Name: "elimisteve", Age: 27, CreatedAt: time.Now()}

	_, err = db.Exec(InsertProfile, &elimi.Name, &elimi.Age)
	maybePanic(err)


	// Find elimisteve in DB
	rows, err := db.Query(SelectProfileByName, "elimisteve")
	maybePanic(err)
	defer rows.Close()

	var steve Profile
	var rowNum int
	// Iterate over all of elimisteve's profiles
	for rows.Next() {
		err = rows.Scan(&steve.Id, &steve.Name, &steve.Age, &steve.CreatedAt)
		maybePanic(err)
		// log.Printf("elimisteve's Profile #%d: %+v\n", rowNum, steve)
		rowNum++
	}
	log.Printf("elimisteve's latest Profile:  %+v\n", steve)


	// Update profile by incrementing age. Notice multiple WHERE clauses
	q := "UPDATE goprofiletest SET Age = $1 WHERE Name = $2 AND Age = $3"
	_, err = db.Exec(q, steve.Age+1, steve.Name, steve.Age)
	maybePanic(err)


	var newSteve Profile
	// Find the profile we found last time and scan it into a new Profile
	err = db.QueryRow(SelectProfileByID, steve.Id).Scan(&newSteve.Id,
		&newSteve.Name, &newSteve.Age, &newSteve.CreatedAt)
	maybePanic(err)
	log.Printf("elimisteve's updated Profile: %+v\n", newSteve)


	// Delete profile
	_, err = db.Exec(DeleteProfileByID, newSteve.Id)
	maybePanic(err)
}

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

type Profile struct {
	Id        int64
	Name      string
	Age       int
	CreatedAt time.Time
	// Interests []string  // TODO
}
