package routers

import (
	"api/src/controllers"
	"net/http"
)

var routersUsers = []Router{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		AuthenticationRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: controllers.SearchUsers,
		AuthenticationRequired: false,
	},
	{
		URI:    "/users/{usersId}",
		Method: http.MethodGet,
		Function: controllers.SearchUser,
		AuthenticationRequired: false,
	},
	{
		URI:    "/users/{usersId}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		AuthenticationRequired: false,
	},
	{
		URI:    "/users/{usersId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		AuthenticationRequired: false,
	},
}