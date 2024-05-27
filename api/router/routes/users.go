package routes

import (
	"net/http"

	"github.com/carlos1377/devbook/api/controllers"
)

var userRoutes = []Route{
	{
		URI:      "/usuarios",
		Method:   http.MethodPost,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		URI:      "/usuarios",
		Method:   http.MethodGet,
		Function: controllers.GetUsers,
		NeedAuth: false,
	},
	{
		URI:      "/usuarios/{id}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		NeedAuth: false,
	},
	{
		URI:      "/usuarios/{id}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: false,
	},
	{
		URI:      "/usuarios/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: false,
	},
}
