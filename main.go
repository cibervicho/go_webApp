package main

import (
	"log"
	"net/http"
	"github.com/cibervicho/go_webApp/movies"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/movies", movies.Index)
	
	http.HandleFunc("/movies/show", movies.Show)
	
	http.HandleFunc("/movies/create", movies.Create)
	http.HandleFunc("/movies/create/process", movies.CreateProcess)
	
	http.HandleFunc("/movies/update", movies.Update)
	http.HandleFunc("/movies/update/process", movies.UpdateProcess)
	
	http.HandleFunc("/movies/delete/process", movies.DeleteProcess)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/movies", http.StatusSeeOther)
}