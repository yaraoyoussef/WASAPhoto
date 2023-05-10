package api

// DONE

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// store values of error and db profile
	var err error
	var profile database.Profile

	// get id of current user
	cUser := extractBearer(r.Header.Get("Authorization"))
	// get id requested by user
	id := ps.ByName("id")

	// check if id is empty
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// check if cUser is banned by other user
	banned, err := rt.db.CheckForBan(id, cUser)
	// handle error
	if err != nil {
		ctx.Logger.WithError(err).Error("an error occured")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// if cUser is banned, do not allow retrieval
	if banned {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// get profile from db
	profile, err1 := rt.db.GetProfile(User{ID: id}.ToDatabase(), User{ID: cUser}.ToDatabase())
	// error handling
	if err1 != nil {
		ctx.Logger.WithError(err).Error("error while executing request")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert to front end profile so that user can visualize it
	var frontendProfile Profile
	frontendProfile.FromDatabase(profile)

	// send profile to user
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(frontendProfile)

}
