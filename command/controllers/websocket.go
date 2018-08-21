package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sysco-middleware/commander-boilerplate/command/common"
	"github.com/sysco-middleware/commander-boilerplate/command/websocket"
)

// ConsumeEvents starts consuming events
// all received events are broadcasted over the websocket network
func ConsumeEvents() {
	events, _ := common.Commander.NewEventsConsumer()

	for {
		event := <-events
		data, err := json.Marshal(event)

		if err != nil {
			continue
		}

		// TODO: events should never be broadcasted
		common.Socket.Broadcast(string(data))
	}
}

// OnWebsocket handles a new websocket request.
// The session is added into the websocket hub and receives new events.
func OnWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	client := &websocket.Client{
		Hub:  common.Socket,
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	client.Hub.Register <- client
	go client.WritePump()
}
