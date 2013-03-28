// Steve Phillips / elimisteve
// 2012.03.31

package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Sbitter struct {
	Username   string
	Password   string
	Age        int
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func maybePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Establish connection to local MongoDB server
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Printf("Install MongoDB then try again!\n")
		return
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Create new *Sbitter objects for elimisteve and swiss
	now := time.Now()
	elimisteve := &Sbitter{"elimisteve", "ilovego!!", 26, now, now}
	swiss := &Sbitter{"swiss", "bananapantsftw", 23, now, now}

	// Save both to the collection of "sbitters" in database "sbitter"
	c := session.DB("sbitter").C("sbitters")
	err = c.Insert(elimisteve, swiss)
	maybePanic(err)

	// Find elimisteve and swiss in DB
	steve := Sbitter{}
	err = c.Find(bson.M{"username": "elimisteve"}).One(&steve)
	maybePanic(err)

	mike := Sbitter{}
	err = c.Find(bson.M{"username": "swiss"}).One(&mike)
	maybePanic(err)

	// Compare the lengths of their... lifetimes
	fmt.Println("Steve is", steve.Age-mike.Age, "years older than Mike")
}
