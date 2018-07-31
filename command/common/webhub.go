package common

import (
	"github.com/gorilla/mux"
	"github.com/sysco-middleware/commander-boilerplate/command/websocket"
)

var (
	// Socket holds the global websocket struct
	Socket *websocket.Hub

	// Router holds the global router struct
	Router *mux.Router
)

// OpenWebHub opens a new socket and router interface
func OpenWebHub() {
	Router = mux.NewRouter()
	Socket = websocket.NewHub()
}
