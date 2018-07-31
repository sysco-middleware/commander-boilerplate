package common

import "github.com/gorilla/mux"

// Router holds the global router struct
var Router *mux.Router

// OpenWebHub opens a new router interface
func OpenWebHub() *mux.Router {
	Router = mux.NewRouter()
	return Router
}
