package main

import (
	"net/http"

	"github.com/sysco-middleware/commander-boilerplate/command/common"
	"github.com/sysco-middleware/commander-boilerplate/command/controllers"
	"github.com/sysco-middleware/commander-boilerplate/command/rest"
)

func main() {
	common.OpenWebHub()
	common.OpenCommander()

	common.Router.HandleFunc("/command/{command}", rest.Use(controllers.OnCommand, Authentication)).Methods("POST")
	common.Router.HandleFunc("/updates", rest.Use(controllers.OnWebsocket, Authentication)).Methods("GET")

	common.Commander.CloseOnSIGTERM()
	go common.Commander.StartConsuming()
	go controllers.ConsumeEvents()

	http.ListenAndServe(":8080", common.Router)
}

// Authentication validates if the given request is authenticated.
// If the request is not authenticated is a 401 returned.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// <- authenticate request and provide the users dataset key
		next.ServeHTTP(w, r)
	}
}
