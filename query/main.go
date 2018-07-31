package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sysco-middleware/commander-boilerplate/query/common"
	"github.com/sysco-middleware/commander-boilerplate/query/controllers"
	"github.com/sysco-middleware/commander-boilerplate/query/rest"
)

func main() {
	common.OpenDatabase()
	router := common.OpenWebHub()

	router.HandleFunc("/find/{id}", rest.Use(controllers.FindByID, Authentication)).Methods("GET")

	http.ListenAndServe(":8080", router)
}

// Authentication validates if the given request is authenticated.
// If the request is not authenticated is a 401 returned.
func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// <- authenticate request and provide the users dataset key
		next.ServeHTTP(w, r)
	}
}
