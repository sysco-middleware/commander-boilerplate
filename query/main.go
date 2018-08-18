package main

import (
	"net/http"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander-boilerplate/query/common"
	"github.com/sysco-middleware/commander-boilerplate/query/controllers"
	"github.com/sysco-middleware/commander-boilerplate/query/rest"
)

func main() {
	common.OpenDatabase()
	router := common.OpenWebHub()

	router.HandleFunc("/find/{id}", rest.Use(controllers.FindByID, Authentication)).Methods("GET")
	router.HandleFunc("/find/", rest.Use(controllers.FindAll, Authentication)).Methods("GET")
	router.HandleFunc("/find/name/last/{lastName}", rest.Use(controllers.FindByLastName, Authentication)).Methods("GET")

	server := &http.Server{
		Addr:         os.Getenv("HOST_ADDRESS"),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}

// Authentication validates if the given request is authenticated.
// If the request is not authenticated is a 401 returned.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// <- authenticate request and provide the users dataset key
		next.ServeHTTP(w, r)
	}
}
