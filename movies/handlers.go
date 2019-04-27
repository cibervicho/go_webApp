package movies

import (
	"github.com/cibervicho/go_webApp/config"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	mve, err := AllMovies()
	if err != nil {
		http.Error(w, http.StatusText(500)+err.Error(), http.StatusInternalServerError)
		return
	}
	
	config.TPL.ExecuteTemplate(w, "movies.gohtml", mve)
}