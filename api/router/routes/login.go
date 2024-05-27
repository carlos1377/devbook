package routes

import (
	"net/http"

	"github.com/carlos1377/devbook/api/controllers"
)

var loginRoute = Route{
	URI:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	NeedAuth: false,
}
