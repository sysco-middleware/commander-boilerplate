package common

import (
	"github.com/gorilla/mux"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/command/websocket"
)

var (
	// Commander holds the global commander struct
	Commander *commander.Commander

	// Socket holds the global websocket struct
	Socket *websocket.Hub

	// Router holds the global router struct
	Router *mux.Router
)
