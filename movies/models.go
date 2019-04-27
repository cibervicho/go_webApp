package movies

import (
	//"errors"
	"github.com/cibervicho/go_webApp/config"
	"github.com/globalsign/mgo/bson"
	//"net/http"
	//"strconv"
)

type Movie struct {
	// add the ID and tags if you need them
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Title string   `json:"title" bson:"title"`
	Year  string   `json:"year" bson:"year"`
	Rated string   `json:"rated" bson:"rated"`
	Genre []string `json:"genre" bson:"genre"`
	Plot  string   `json:"plot" bson:"plot"`
}

func AllMovies() ([]Movie, error) {
	mve := []Movie{}
	err := config.IMBD.Find(nil).Sort("year").All(&mve)
	if err != nil {
		return nil, err
	}
	return mve, nil
}