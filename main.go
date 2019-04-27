package main

import (
	"log"
	"net/http"
	"github.com/cibervicho/go_webApp/movies"
)

/*type person struct {
	FirstName string
	Age int
}*/

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/movies", movies.Index)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/movies", http.StatusSeeOther)
}