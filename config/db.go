package config

import (
	"fmt"
	"github.com/globalsign/mgo"
	"log"
)

const (
	url = "localhost:27017"
	my_db = "moviesdb"
	my_coll = "imdb"
)

// database
var DB *mgo.Database

// collections
var IMBD *mgo.Collection

func init() {
	// connecting to the mongodb server
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	//defer session.Close()

	// make sure we can ping the server
	if err = session.Ping(); err != nil {
		panic(err)
	}

	DB = session.DB(my_db)
	IMBD = DB.C(my_coll)
	
	fmt.Println("Successfully connected to mongodb server at", url)
}