package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

var CurrentUser User

// function used to change username of current user
func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// take id
	id := ps.ByName("id")

	// extract from header
	userReq := extractBearer(r.Header.Get("Authorization"))
	// validation
	valid := validateUser(id, userReq)
	if valid != 0 {
		w.WriteHeader(valid)
		return
	}

	// decode the new username
	var newUsername Username
	err := json.NewDecoder(r.Body).Decode(&newUsername)
	if err != nil {
		// body was not parseable JSON, rejected
		ctx.Logger.WithError(err).Error("error parsing json file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if username format is valid
	if !IsValid(newUsername.Username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.SetUsername(User{ID: id}.ToDatabase(), newUsername.ToDatabase())
	if errors.Is(err, database.ErrCouldNotModify) {
		w.WriteHeader(http.StatusConflict)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("couldn't update username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// response
	w.WriteHeader(http.StatusNoContent)
}
