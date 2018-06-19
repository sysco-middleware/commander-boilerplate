package websocket

import (
	"encoding/json"

	"github.com/sysco-middleware/commander"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	Clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client

	Commander *commander.Commander
}

// NewHub create a new hub
func NewHub(commander *commander.Commander) *Hub {
	hub := &Hub{
		broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Commander:  commander,
	}

	go hub.consume()
	go hub.run()

	return hub
}

func (h *Hub) consume() {
	consumer := h.Commander.Consume("events")

	for msg := range consumer.Messages {
		event := commander.Event{}
		json.Unmarshal(msg.Value, &event)

		data, err := json.Marshal(event)

		if err != nil {
			continue
		}

		h.Broadcast(string(data))
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client] = true
		case client := <-h.Unregister:
			if _, ok := h.Clients[client]; ok {
				delete(h.Clients, client)
				close(client.Send)
			}
		case message := <-h.broadcast:
			for client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, client)
				}
			}
		}
	}
}

// Broadcast a new message to all connected clients
func (h *Hub) Broadcast(message string) {
	h.broadcast <- []byte(message)
}
