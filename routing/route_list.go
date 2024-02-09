package routing

import (
	"event_ticket/handler"
	"net/http"
)

var Routes = []Route{
	{
		Method:  http.MethodGet,
		Path:    "/home",
		Handler: handler.Home,
	},
	{
		Method:  http.MethodPost,
		Path:    "/buy",
		Handler: handler.Buy,
	},
}
