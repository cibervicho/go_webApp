package movies

import (
	"errors"
	"github.com/cibervicho/go_webApp/config"
	"github.com/globalsign/mgo/bson"
	"net/http"
)

type Movie struct {
	// add the ID and tags if you need them
	ID    bson.ObjectId `json:"id" bson:"_id"`
	My_id string
	Title string   `json:"title" bson:"title"`
	Year  string   `json:"year" bson:"year"`
	Rated string   `json:"rated" bson:"rated"`
	Genre []string `json:"genre" bson:"genre"`
	Plot  string   `json:"plot" bson:"plot"`
}

type InsertMovie struct {
	// add the ID and tags if you need them
	//ID    string `json:"id" bson:"_id"`
	Title string   `json:"title" bson:"title"`
	Year  string   `json:"year" bson:"year"`
	Rated string   `json:"rated" bson:"rated"`
	Genre []string `json:"genre" bson:"genre"`
	Plot  string   `json:"plot" bson:"plot"`
}

func AllMovies() ([]Movie, error) {
	mve := []Movie{}
	err := config.IMBD.Find(bson.M{}).Sort("year").All(&mve)
	if err != nil {
		return nil, err
	}
	// [hack] saving the string representation of mongo's _id
	for i, _ := range mve {
		mve[i].My_id = mve[i].ID.Hex()
	}

	return mve, nil
}

func OneMovie(r *http.Request) (Movie, error) {
	mv := Movie{}
	id := r.FormValue("my_id")

	if id == "" {
		return mv, errors.New("400. Bad Request.")
	}
	err := config.IMBD.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&mv)
	if err != nil {
		return mv, err
	}
	// [hack] saving the string representation of mongo's _id
	mv.My_id = mv.ID.Hex()
	
	return mv, nil
}

func PutMovie(r *http.Request) (InsertMovie, error) {
	// get form values
	mv := InsertMovie{}
	
	mv.Title = r.FormValue("title")
	mv.Year = r.FormValue("year")
	mv.Rated = r.FormValue("rated")
	mv.Genre = append(mv.Genre, r.FormValue("genre"))
	mv.Plot = r.FormValue("plot")

	// validate form values
	if mv.Year == "" || mv.Title == "" {
		return mv, errors.New("400. Bad request. Title and Year fields must be provided.")
	}

	// insert values
	err := config.IMBD.Insert(mv)
	if err != nil {
		return mv, errors.New("500. Internal Server Error." + err.Error())
	}
	return mv, nil
}

func UpdateMovie(r *http.Request) (Movie, error) {
	// get form values
	mv := Movie{}
	
	mv.ID = bson.ObjectIdHex(r.FormValue("my_id"))
	mv.Title = r.FormValue("title")
	mv.Year = r.FormValue("year")
	mv.Rated = r.FormValue("rated")
	mv.Genre = append(mv.Genre, r.FormValue("genre"))
	mv.Plot = r.FormValue("plot")

	// validate form values
	if mv.Year == "" || mv.Title == "" {
		return mv, errors.New("400. Bad request. Title and Year fields must be provided.")
	}

	// update values
	err := config.IMBD.Update(bson.M{"_id": mv.ID}, &mv)
	if err != nil {
		return mv, err
	}
	return mv, nil
}

func DeleteMovie(r *http.Request) error {
	id := r.FormValue("my_id")
	if id == "" {
		return errors.New("400. Bad Request.")
	}

	err := config.IMBD.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}