package common

import (
	"github.com/gorilla/mux"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/command/websocket"
)

var (
	// Commander stores the global commander struct
	Commander *commander.Commander

	// Socket stores the global websocket struct
	Socket *websocket.Hub

	// Router stores the global router struct
	Router *mux.Router
)
