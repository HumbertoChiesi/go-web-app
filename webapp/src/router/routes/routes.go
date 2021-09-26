package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Rout represents all the routes of the Web App
type Rout struct {
	URI                string
	Method             string
	Function           func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

//Config inserts all the routes in the router
func Config(r *mux.Router) *mux.Router {
	routs := loginRouts
	routs = append(routs, usersRouts...)

	for _, rout := range routs {
		r.HandleFunc(rout.URI, rout.Function).Methods(rout.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
