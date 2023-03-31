package api

// TO FINISH

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

var CurrentUser User

// function used to change username of current user
func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// take id of user (but in api we take username in path only...)
	id, err := strconv.Atoi(ps.ByName("ID"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// take username of user, and if it is not equal to the current username, display error (we cannot change another user's username)
	username := r.URL.Query().Get("username")
	if username != CurrentUser.Username {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var newUsername string
	err = json.NewDecoder(r.Body).Decode(&newUsername)
	if err != nil {
		// body was not parseable JSON, rejected
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var updatedUser User
	updatedUser.Username = newUsername
	if !updatedUser.IsValid(updatedUser.Username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedUser.ID = strconv.Itoa(int(id))

	err = rt.db.SetUsername(updatedUser.ToDatabase())
	if errors.Is(err, database.ErrCouldNotModify) {
		w.WriteHeader(http.StatusConflict)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).WithField("ID", id).Error("couldn't update username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
