package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/sysco-middleware/commander-boilerplate/projector/models"
	"github.com/sysco-middleware/commander-boilerplate/query/common"
	"github.com/sysco-middleware/commander-boilerplate/query/rest"
)

// FindByID finds a user by the given id
func FindByID(w http.ResponseWriter, r *http.Request) {
	res := rest.Response{ResponseWriter: w}
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])

	if err != nil {
		res.SendPanic(err.Error(), nil)
		return
	}

	user := models.Users{}
	query := common.Database.Where(models.Users{ID: &id}).First(&user)

	if query.Error == gorm.ErrRecordNotFound {
		res.SendNotFound()
		return
	}

	if query.Error != nil {
		res.SendPanic(query.Error.Error(), nil)
		return
	}

	res.SendOK(user)
}
