package main

import (
	"net/http"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/command/common"
	"github.com/sysco-middleware/commander-boilerplate/command/controllers"
	"github.com/sysco-middleware/commander-boilerplate/command/rest"
	"github.com/sysco-middleware/commander-boilerplate/command/websocket"
)

func main() {
	servers := os.Getenv("KAFKA_SERVERS")
	group := os.Getenv("KAFKA_GROUP")

	producer := commander.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":    servers,
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})

	consumer := commander.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    servers,
		"group.id":             group,
		"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})

	common.Router = mux.NewRouter()
	common.Socket = websocket.NewHub()

	common.Commander = &commander.Commander{
		Consumer: consumer,
		Producer: producer,
	}

	common.Router.HandleFunc("/command/{command}", rest.Use(controllers.OnCommand, Authentication)).Methods("POST")
	common.Router.HandleFunc("/updates", rest.Use(controllers.OnWebsocket, Authentication)).Methods("GET")

	common.Commander.CloseOnSIGTERM()
	common.Commander.StartConsuming()

	go controllers.ConsumeEvents()
	http.ListenAndServe(":8080", common.Router)
}

// Authentication validates if the given request is authenticated.
// If the request is not authenticated is a 401 returned.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// <- authenticate request and provide the users dataset key
		next.ServeHTTP(w, r)
	}
}
