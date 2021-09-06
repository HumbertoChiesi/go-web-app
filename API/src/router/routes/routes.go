package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

//Rout represents all the routes from the API
type Rout struct {
	URI                string
	Method             string
	Function           func(http.ResponseWriter, *http.Request)
	NeedAuthentication bool
}

//Config inserts all the routes in the router
func Config(r *mux.Router) *mux.Router {
	routs := usersRoutes
	routs = append(routs, loginRout)

	for _, rout := range routs {
		if rout.NeedAuthentication {
			r.HandleFunc(rout.URI,
				middlewares.Logger(middlewares.Authenticate(rout.Function)),
			).Methods(rout.Method)
		} else {
			r.HandleFunc(rout.URI, middlewares.Logger(rout.Function)).Methods(rout.Method)
		}
	}
	return r
}
