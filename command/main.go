package main

import (
	"net/http"

	"github.com/sysco-middleware/commander-boilerplate/command/common"
	"github.com/sysco-middleware/commander-boilerplate/command/controllers"
	"github.com/sysco-middleware/commander-boilerplate/command/rest"
)

func main() {
	router, _ := common.OpenWebHub()
	commander := common.OpenCommander()

	router.HandleFunc("/command/{command}", rest.Use(controllers.OnCommand, Authentication)).Methods("POST")
	router.HandleFunc("/updates", rest.Use(controllers.OnWebsocket, Authentication)).Methods("GET")

	commander.CloseOnSIGTERM()
	go commander.StartConsuming()
	go controllers.ConsumeEvents()

	server := &http.Server{
		Addr:         os.Getenv("HOST_ADDRESS"),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}

// Authentication validates if the given request is authenticated.
// If the request is not authenticated is a 401 returned.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// <- authenticate request and provide the users dataset key
		next.ServeHTTP(w, r)
	}
}
