// Steve Phillips / elimisteve
// 2013.10.27

package types

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"os"
	"time"
)

const (
	createTestTable = `CREATE TABLE goprofiletest (
	Id        serial PRIMARY KEY NOT NULL,
	Name      varchar(100) NOT NULL,
	Age       integer NOT NULL,
	CreatedAt timestamp NOT NULL DEFAULT current_timestamp
);`

	insertProfile = "INSERT INTO goprofiletest (Id, Name, Age, CreatedAt) VALUES (DEFAULT, $1, $2, DEFAULT)"
)

var (
	db *sql.DB

	postgresConnStr = "user=postgres dbname=postgres host=localhost sslmode=disable password=" + os.Getenv("PGPASS")
)

func init() {
	if db == nil {
		var e error
		db, e = sql.Open("postgres", postgresConnStr)
		if e != nil {
			log.Fatal(e)
		}
		log.Printf("`db` value set\n")
	}
}

type Profile struct {
	Id        int64
	Name      string
	Age       int
	CreatedAt time.Time
}

func (p *Profile) InsertFields() []interface{} {
	return []interface{}{
		&p.Name,
		&p.Age,
	}
}

func (p *Profile) Save() error {
	_, err := db.Exec(insertProfile, p.InsertFields()...)
	return err
}
