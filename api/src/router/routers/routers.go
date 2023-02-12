package routers

import "net/http"


//Router represents all routes from the API
type Router struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	AuthenticationRequired bool
}