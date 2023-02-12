package router

import "github.com/gorilla/mux"

//Generate will return a router with the configured routers
func Generate() *mux.Router {
	return mux.NewRouter()
}