package hub

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/command/hub/rest"
	"github.com/sysco-middleware/commander-boilerplate/command/hub/websocket"
)

// Hub create a new hub that starts a http server and websocket server
type Hub struct {
	Socket    *websocket.Hub
	Router    *mux.Router
	commander *commander.Commander
}

// NewHub create a new hub with the given commander instance
func NewHub(commander *commander.Commander) *Hub {
	socket := websocket.NewHub(commander)
	router := mux.NewRouter()

	hub := &Hub{
		socket,
		router,
		commander,
	}

	return hub
}

// Open open the http servers
func (h *Hub) Open() {
	http.ListenAndServe(":8080", h.Router)
}

// HandleCommandRequest handle a new http command request
func (h *Hub) HandleCommandRequest(w http.ResponseWriter, r *http.Request) {
	res := rest.Response{ResponseWriter: w}
	params := r.URL.Query()
	vars := mux.Vars(r)

	sync, _ := strconv.ParseBool(params.Get("sync"))
	body, _ := ioutil.ReadAll(r.Body)

	action := vars["command"]
	command := commander.NewCommand(action, body)

	if sync {
		event, err := h.commander.SyncCommand(command)

		if err != nil {
			res.SendPanic(err.Error(), command)
			return
		}

		res.SendOK(event)
		return
	}

	err := h.commander.AsyncCommand(command)

	if err != nil {
		res.SendPanic(err.Error(), nil)
		return
	}

	res.SendCreated(command)
}

// HandleWebsocketRequest handle a new websocket request
func (h *Hub) HandleWebsocketRequest(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	client := &websocket.Client{
		Hub:  h.Socket,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
}
