package routes

import (
	"net/http"
	"webapp/src/middlewares"

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
	routs = append(routs, homePageRout)
	routs = append(routs, postsRout...)

	for _, rout := range routs {
		if rout.NeedAuthentication {
			r.HandleFunc(rout.URI,
				middlewares.Logger(middlewares.Authenticate(rout.Function)),
			).Methods(rout.Method)
		} else {
			r.HandleFunc(rout.URI,
				middlewares.Logger(rout.Function),
			).Methods(rout.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
