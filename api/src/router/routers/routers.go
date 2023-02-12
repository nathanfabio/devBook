package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Router represents all routes from the API
type Router struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}


//Setup Puts all routes inside the router
func Setup(r *mux.Router) *mux.Router {
	routers := routersUsers

	for _, router := range routers {
		r.HandleFunc(router.URI, router.Function).Methods((router.Method))
	}

	return r
}