package router

import (
	"api/src/router/routers"

	"github.com/gorilla/mux"
)

//Generate will return a router with the configured routers
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routers.Setup(r)
}