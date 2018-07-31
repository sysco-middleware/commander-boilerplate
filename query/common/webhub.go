package common

import "github.com/gorilla/mux"

// Router holds the global router struct
var Router *mux.Router

// OpenWebHub opens a new router interface
func OpenWebHub() {
	Router = mux.NewRouter()
}
