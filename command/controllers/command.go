package controllers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sysco-middleware/commander"
	"github.com/sysco-middleware/commander-boilerplate/command/common"
	"github.com/sysco-middleware/commander-boilerplate/command/rest"
)

// OnCommand handles a new command request.
// The received command can be executed in a sync or async manner.
func OnCommand(w http.ResponseWriter, r *http.Request) {
	res := rest.Response{ResponseWriter: w}
	params := r.URL.Query()
	vars := mux.Vars(r)

	sync, _ := strconv.ParseBool(params.Get("sync"))
	body, _ := ioutil.ReadAll(r.Body)

	action := vars["command"]
	command := commander.NewCommand(action, body)

	if sync {
		event, err := common.Commander.SyncCommand(command)

		if err != nil {
			res.SendPanic(err.Error(), command)
			return
		}

		if !event.Acknowledged {
			res.SendPanic(event.Action, event)
			return
		}

		res.SendOK(event)
		return
	}

	err := common.Commander.AsyncCommand(command)

	if err != nil {
		res.SendPanic(err.Error(), nil)
		return
	}

	res.SendCreated(command)
}
