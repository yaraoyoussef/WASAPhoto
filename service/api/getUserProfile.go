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
	// get username requested by user
	username := r.URL.Query().Get("username")
	// check if username is empty
	if username == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	} else {
		// get profile from db
		profile, err = rt.db.GetProfile(username)
	}

	if err != nil {
		// we have an internal server error
		ctx.Logger.WithError(err).Error("error, please try again later")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// convert to front end profile so that user can visualize it
	var frontendProfile Profile
	frontendProfile.FromDatabase(profile)

	// send profile to user
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(frontendProfile)

}
