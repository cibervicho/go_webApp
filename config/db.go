package config

import (
	"fmt"
	"github.com/globalsign/mgo"
	"log"
	"os"
)

const (
	//url = "localhost:27017"
	//url = "mongodb://192.168.1.66:27017/moviesdb"}
	//url = "laptop-03garp8v"
	my_db = "moviesdb"
	my_coll = "imdb"
)

var url string

// database
var DB *mgo.Database

// collections
var IMBD *mgo.Collection

func init() {
	if 1 == len(os.Args) {
		url = "localhost:27017"
		log.Printf("No db address specified, using '%v' for now", url)
	} else {
		// set the specified mongodb database in command line
		url = os.Args[1]
		log.Printf("Using the IP '%v' as the db address", url)
	}
	
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