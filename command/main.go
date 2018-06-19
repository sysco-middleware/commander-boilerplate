package main

import (
	"net/http"
	"os"

	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/command/hub"
	"github.com/sysco-middleware/commander-boilerplate/command/hub/rest"
)

var cmd *commander.Commander

func main() {
	host := os.Getenv("KAFKA_HOST")
	group := os.Getenv("KAFKA_GROUP")

	cmd = &commander.Commander{
		Producer: commander.NewProducer(host),
		Consumer: commander.NewConsumer(host, group),
	}

	go cmd.ReadMessages()
	go cmd.CloseOnSIGTERM()

	// Initialize a new hub
	hub := hub.NewHub(cmd)
	hub.Router.HandleFunc("/command/{command}", rest.Use(hub.HandleCommandRequest, authentication)).Methods("POST")
	hub.Router.HandleFunc("/updates", rest.Use(hub.HandleWebsocketRequest, authentication)).Methods("GET")

	hub.Open()
}

func authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// <- authenticate the user
		next.ServeHTTP(w, r)
	})
}
