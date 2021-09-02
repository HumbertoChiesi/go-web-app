package router

import "github.com/gorilla/mux"

//Generate will return a router with the configured routes
func Generate() *mux.Router {
	return mux.NewRouter()
}
