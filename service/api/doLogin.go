package api

// DONE

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// function that handles user's Login
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// read content of the user from request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		// body was not parsable, so it gets rejected
		w.WriteHeader(http.StatusBadRequest)
		return
		// validation of user's structure content
	} else if !user.IsValid(user.Username) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// send user info to db
	dbUser, err := rt.db.Login(user.ToDatabase())
	if err != nil {
		// handle error on our side. Log it and send 500 to user
		// 500 code error should be in api operation responses??
		ctx.Logger.WithError(err).Error("cannot login user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// over-writing user from info in db
	user.FromDatabase(dbUser)

	// send output to user
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(user)
}
